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
	dbname   = "project"
)

type Detials struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Position string  `json:"position"`
	Salary   float64 `json:"salary"`
	Mail     string  `json:"mail"`
}

var (
	detials = []Detials{}
)

func main() {
	var err error
	connect := fmt.Sprintf("host = %s port = %d user = %s password = %s dbname = %s sslmode = disable ", host, port, user, password, dbname)
	db, err = sql.Open("postgres", connect)
	if err != nil {
		log.Fatal("error connecting database !")
		return
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal("error connecting ping database")
		return
	}
	fmt.Println("successfully connecting database")

	createtable()

	g := gin.Default()
	g.Use(gin.Logger())
	g.GET("/employe", getalldetials)
	g.GET("/employe/:id", getdetials_byid)
	g.POST("/employe", createdetiales)
	g.PUT("/employe/:id", updatedetials)
	g.DELETE("/employe/:id", delete_detials)

	g.Run(":5050")
	fmt.Println("starting server at :1010")

}

func createtable() {
	query := `CREATE TABLE IF NOT EXISTS employee 
	( 
	id SERIAL PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	position VARCHAR(255) NOT NULL,
	salary INT NOT NULL,
	mail VARCHAR(255) NOT NULL
	);`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("error creating table", err)
		return
	}
	fmt.Println("successfully creating table")
}
func createdetiales(c *gin.Context) {
	var newemploye Detials
	err := c.ShouldBindJSON(&newemploye)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error inserting the data"})
		return
	}

	_, err = db.Exec("INSERT INTO employee(id,name,position,salary,mail) VALUES ($1,$2,$3,$4,$5)", newemploye.Id, newemploye.Name, newemploye.Position, newemploye.Salary, newemploye.Mail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error insrting the data to table"})
		return
	}
	//detials = append(detials, newemploye)
	c.JSON(http.StatusOK, newemploye)

}

func getalldetials(c *gin.Context) {
	rows, err := db.Query("SELECT id, name, position, salary, mail FROM employee")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error getting the employe detials"})
		return
	}
	defer rows.Close()
	var employes []Detials
	for rows.Next() {
		var employe Detials
		err := rows.Scan(&employe.Id, &employe.Name, &employe.Position, &employe.Salary, &employe.Mail)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// employe := map[string]interface{}{
		// 	"Id":       id,
		// 	"Name":     name,
		// 	"Position": position,
		// 	"Salary":   salary,
		// 	"Mail":     mail,
		// }
		employes = append(employes, employe)

	}
	c.JSON(http.StatusOK, employes)
}
func getdetials_byid(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employe id not found"})
		return
	}
	var employe Detials
	err = db.QueryRow("SELECT id, name, position, salary, mail FROM employee WHERE id=$1", id).Scan(&employe.Id, &employe.Name, &employe.Position, &employe.Salary, &employe.Mail)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "employe detials not found"})
		return
	}
	c.JSON(http.StatusOK, employe)
}
func updatedetials(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employe id not found"})
		return
	}
	var updateemploye Detials
	if err := c.ShouldBindJSON(&updateemploye); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err = db.Exec("UPDATE employee SET name=$1,position=$2,salary=$3,mail=$4 WHERE id=$5", updateemploye.Name, updateemploye.Position, updateemploye.Salary, updateemploye.Mail, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error updating detials in id"})
		return
	}
	updateemploye.Id = id
	c.JSON(http.StatusOK, updateemploye)
}
func delete_detials(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id not found"})
		return
	}
	_, err = db.Exec("DELETE FROM employee  WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successfully delete "})
}
