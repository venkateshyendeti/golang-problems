package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Detials struct {
	Bookname string `json:"bookname"`
	Author   string `json:"author"`
	Price    int64  `json:"price"`
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Venkey83$"
	db_name  = "student"
)

var db *sql.DB

func main() {
	var err error
	connect := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, db_name)
	db, err = sql.Open("postgres", connect)
	if err != nil {
		log.Fatal("error conecting database")
		return
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal("error connecting ping database", err)
		return
	}
	fmt.Println("successfully connecting database")

	create_Table()

	g := gin.Default()
	g.Use(gin.Logger())
	g.GET("/book", getbooks)
	g.POST("/book", createbook)
	g.GET("/book/:bookname", getbookbyname)
	g.PUT("/book/:bookname", updatebook)
	g.DELETE("/book/:bookname", deletebook)
	g.Run(":2020")

}
func create_Table() {
	query := `CREATE TABLE IF NOT EXISTS book(
	Bookname VARCHAR(250) NOT NULL,
	Author VARCHAR(250) NOT NULL,
    Price INT NOT NULL
	);`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("error creating table")
		return
	}
	fmt.Println("successfully creating table")

}
func getbooks(c *gin.Context) {
	rows, err := db.Query("SELECT  bookname,author,price FROM book")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var books []Detials
	for rows.Next() {
		var book Detials
		if err := rows.Scan(&book.Bookname, &book.Author, &book.Price); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error scanning"})
			return
		}
		books = append(books, book)
	}

	c.JSON(http.StatusOK, books)
}
func createbook(c *gin.Context) {
	var newBook []Detials
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// for _, item := range newbook {
	// 	_, err := db.Exec("INSERT INTO book (bookname, author, price) VALUES ($1,$2,$3)", item.Bookname, item.Author, item.Price)
	// 	if err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": "error insert values in the table"})
	// 		return
	// 	}
	// }
	for _, item := range newBook {
		_, err := db.Exec("INSERT INTO book (bookname, author, price) VALUES ($1, $2, $3)",
			item.Bookname, item.Author, item.Price)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error inserting values into the table"})
			return
		}
	}
	//detials = append(detials, newBook...)
	c.JSON(http.StatusOK, newBook)
}
func getbookbyname(c *gin.Context) {
	name := strings.TrimSpace(c.Param("bookname"))
	if name == " " {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name not found"})
		return
	}
	// for _, item := range detials {
	// 	if item.Bookename==name {
	// 		c.JSON(http.StatusOK,item)
	// 		return
	// 	}
	//}
	var book Detials
	err := db.QueryRow("SELECT bookname, author, price FROM book WHERE bookname=$1", name).Scan(&book.Bookname, &book.Author, &book.Price)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "book detials not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, book)
}
func updatebook(c *gin.Context) {
	name := strings.TrimSpace(c.Param("bookname"))
	if name == " " {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name not found "})
		return
	}
	var updatebook Detials
	if err := c.ShouldBindJSON(&updatebook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	update := `UPDATE book SET Bookname=$1,Author=$2,price=$3 WHERE bookname=$4`
	_, err := db.Exec(update, updatebook.Bookname, updatebook.Author, updatebook.Price, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// for _, item := range detials {
	// 	if item.Bookname == name {
	// 		updatebook.Bookname = name
	// 		c.JSON(http.StatusOK, updatebook)
	// 		//return
	// 	}
	// }

	c.JSON(http.StatusOK, updatebook)
}
func deletebook(c *gin.Context) {
	name := strings.TrimSpace(c.Param("bookname"))
	if name == " " {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name not found"})
		return
	}
	_, err := db.Exec("DELETE FROM book WHERE bookname=$1", name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "book successfully delete"})
}
