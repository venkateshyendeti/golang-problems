package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Venkey83$"
	dbname   = "student"
)

type Detials struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Subject string `json:"subject"`
}

var (
	detials = []Detials{}
)

func main() {
	connect := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	var err error
	db, err = sql.Open("postgres", connect)
	if err != nil {
		log.Fatal("error conecting database", err)
		return
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal("error connecting ping database", err)
		return
	}
	fmt.Println("successfully connected database")

	createtable()

	g := gin.Default()
	g.Use(gin.Logger())
	g.GET("/student", getdetials)
	g.POST("/student", createstudent)
	g.GET("/student/:id", get_byID)
	g.PUT("/student/:id", updatedstudent)
	g.DELETE("/student/:id", deletedstudent)
	g.Run(":4040")
	fmt.Println("server starting :1010")
}
func createtable() {
	query := `
	CREATE TABLE IF NOT EXISTS teacher(
	id SERIAL PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	subject VARCHAR(255) NOT NULL
	);`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("error creating table ", err)
		return
	}
	fmt.Println("successfully creating table")
}

func getdetials(c *gin.Context) {
	_, err := db.Exec("SELECT * FROM teacher")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, detials)
}

func createstudent(c *gin.Context) {
	var newstudent []Detials
	if err := c.ShouldBindJSON(&newstudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, student := range newstudent {

		_, err := db.Exec("INSERT INTO teacher(id,name,subject) VALUES($1,$2,$3)", student.Id, student.Name, student.Subject)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error inserting values in table"})
			return
		}
	}

	detials = append(detials, newstudent...)
	c.JSON(http.StatusOK, newstudent)
}
func get_byID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id not found"})
		return
	}

	var student Detials
	err = db.QueryRow("SELECT id, name, subject FROM teacher WHERE id=$1", id).
		Scan(&student.Id, &student.Name, &student.Subject)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "student not found"})
		return
	}

	c.JSON(http.StatusOK, student)
}
func updatedstudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id not found to the table"})
		return
	}
	var updatedstudent Detials
	if err = c.ShouldBindJSON(&updatedstudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	update := `UPDATE teacher SET name=$1,subject=$2 WHERE id=$3`
	_, err = db.Exec(update, updatedstudent.Name, updatedstudent.Subject, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error updating"})
		return
	}
	updatedstudent.Id = id
	c.JSON(http.StatusOK, updatedstudent)

}
func deletedstudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id not found"})
		return
	}
	// for i, item := range detials {
	// 	if item.Id == id {
	// 		detials = append(detials[:i], detials[i+1:]...)
	// 		c.JSON(http.StatusOK, gin.H{"message": "id successfully deleted"})
	// 		return
	// 	}
	// }
	_, err = db.Exec("DELETE FROM teacher WHERE ID=$1", id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successfully delete record"})

}
