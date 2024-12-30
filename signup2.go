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
	server_host     = "localhost"
	sever_port      = 5432
	server_user     = "postgres"
	server_password = "Venkey83$"
	server_dbname   = "student"
)

var db *sql.DB

type student struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Markes int    `json:"markes"`
}

func main() {
	var err error
	connect := fmt.Sprintf("server_host=%s server_port=%d server_user=%s server_password=%s server_dbname=%s sslmode=disable", server_host, sever_port, server_user, server_password, server_dbname)
	db, err = sql.Open("postgres", connect)
	if err != nil {
		log.Fatal("error connecting database", err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal("error pinging connection", err)
	}
	fmt.Printf("successfully connecting database ")
	createtable()
	r := mux.NewRouter()
	r.HandleFunc("/student", createstudent).Methods("POST")
	fmt.Println("server starting at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
func createtable() {
	query := `CREATE TABLE IF NOT EXCISTS STUDENT(
		id serial PRIMARY KEY,
		username varchar(50) UNIQUE NOT NULL,
		password varchar(50) NOT NULL
	);`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("error creating table", err)
	}
}
func createstudent(w http.ResponseWriter, r *http.Request) {
	var Student []student
	if err := json.NewDecoder(r.Body).Decode(&Student); err != nil {
		log.Fatal("invalid request payloading", http.StatusBadRequest)
		return
	}
	for _, item := range Student {
		_, err := db.Exec("INSERT INTO student(id,name,markes)values($1,$2,$3)", item.Id, item.Name, item.Markes)
		if err != nil {
			http.Error(w, "error creating user", http.StatusInternalServerError)
			log.Fatal("error inserting user", err)
			return
		}

	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "successfully created user"})
}
