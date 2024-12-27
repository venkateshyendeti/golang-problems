package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	dbhost     = "localhost"
	dbport     = 5432
	dbuser     = "postgres"
	dbpassword = "Venkey83$"
	dbname     = "student"
)

type Users struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Markes int    `json:"markes"`
}

var db *sql.DB

func main() {
	var err error
	connect := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbhost, dbport, dbuser, dbpassword, dbname)
	db, err = sql.Open("postgres", connect)
	if err != nil {
		log.Fatal("error connecting to db", err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal("error pinging db", err)
	}
	fmt.Println("connected to db")
	create_Table()

	r := mux.NewRouter()
	//create_Table()
	r.HandleFunc("/users", getusers).Methods("GET")
	r.HandleFunc("/users", createuser).Methods("POST")
	fmt.Println("server starting at :9090")
	log.Fatal(http.ListenAndServe(":9090", r))
}
func create_Table() {
	query := `CREATE TABLE IF NOT EXISTS users(
		id SERIAL PRIMARY KEY,
		name VARCHAR(255),
		markes INT
	);`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("error creating table", err)
	}
	fmt.Println("table created successfully")
}
func getusers(w http.ResponseWriter, r *http.Request) {
	var users []Users
	query := "SELECT * FROM users"
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "error fetching users", http.StatusInternalServerError)
		log.Fatal("error fetching users", err)
		return
	}
	//json.NewEncoder(w).Encode(User)
	defer rows.Close()

	// Scan rows into users slice
	for rows.Next() {
		var user Users
		err := rows.Scan(&user.Id, &user.Name, &user.Markes)
		if err != nil {
			http.Error(w, "error scanning users", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	json.NewEncoder(w).Encode(users)
}
func createuser(w http.ResponseWriter, r *http.Request) {
	var user []Users
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}
	for _, item := range user {

		_, err := db.Exec("INSERT INTO users(id,name,markes) VALUES($1,$2,$3)", item.Id, item.Name, item.Markes)
		if err != nil {
			http.Error(w, "error creating user", http.StatusInternalServerError)
			log.Fatal("error creating user", err)
			return
		}
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "user created successfully"})

}
