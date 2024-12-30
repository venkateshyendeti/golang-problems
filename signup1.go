// package main

// import (
// 	"database/sql"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/mux"
// 	_ "github.com/lib/pq"
// )

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "Venkey83$"
// 	dbname   = "detials"
// )

// var db *sql.DB

// type User struct {
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// }

// // var users =[]User{}
// func main() {
// 	//var db *sql.DB
// 	var err error
// 	connect := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
// 	db, err = sql.Open("postgres", connect)
// 	if err != nil {
// 		log.Fatal("error connecting database", err)

// 	}
// 	defer db.Close()

// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatal("error is pinging connecting ", err)
// 	}

// 	fmt.Println("successfully connecting database")

// 	createtable()
// 	r := mux.NewRouter()
// 	r.HandleFunc("/signup", signuphandler).Methods("POST")
// 	log.Println("starting server at :8080")
// 	log.Fatal(http.ListenAndServe(":8080", r))

// }
// func createtable() {
// 	query := `CREATE TABLE IF NOT EXISTS userr(
// 	id serial PRIMARY KEY,
// 	USERNAME varchar(50) UNIQUE NOT NULL,
// 	PASSWORD varchar(50) NOT NULL
// 	);`
// 	_, err := db.Exec(query)
// 	if err != nil {
// 		log.Fatal("error creating table", http.StatusBadRequest)
// 		return
// 	}
// }
// func signuphandler(w http.ResponseWriter, r *http.Request) {
// 	var User []User
// 	if err := json.NewDecoder(r.Body).Decode(&User); err != nil {
// 		log.Fatal("invalid request payload", http.StatusBadRequest)
// 		return
// 	}
// 	// if User.Username == "" || User.Password == "" {
// 	// 	http.Error(w, "user name and password are required ", http.StatusBadRequest)
// 	// 	return
// 	// }
// 	for _, item := range User {

// 		_, err := db.Exec("INSERT INTO userr(username,password)VALUES ($1,$2)", item.Username, item.Password)
// 		if err != nil {
// 			http.Error(w, "error creating user", http.StatusInternalServerError)
// 			log.Fatal("error inserting error", err)
// 			return
// 		}
// 	}
// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(map[string]string{"message": "user created successfully"})
// }
