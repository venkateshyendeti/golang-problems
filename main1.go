package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Student struct {
	Name   string `json:"name"`
	Class  string `json:"class"`
	Markes int    `json:"markes"`
}

var student = []Student{}

func main() {
	router := gin.Default()
	router.Use(gin.Logger())
	router.GET("/student", getdetials)
	router.GET("/student/:name", get_iddetials)
	router.POST("/student", createstudent)
	router.PUT("/student/:name", updatestudent)
	router.DELETE("/student/:name", delete_student)

	router.Run(":2020")
}
func getdetials(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, student)
}
func get_iddetials(c *gin.Context) {
	name := c.Param("name")
	if name == " " {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name not found"})
		return
	}
	for _, detials := range student {
		if detials.Name == name {
			c.IndentedJSON(http.StatusOK, detials)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "student not found"})
}
func createstudent(c *gin.Context) {
	var newstudent []Student
	if err := c.BindJSON(&newstudent); err != nil {
		return
	}
	student = append(student, newstudent...)
	c.IndentedJSON(http.StatusOK, newstudent)
}
func updatestudent(c *gin.Context) {
	name := c.Param("name")
	if name == " " {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name not found"})
		return
	}
	var updatestudent Student
	if err := c.BindJSON(&updatestudent); err != nil {
		return
	}
	for i, detial := range student {
		if detial.Name == name {
			student[i] = updatestudent
			c.IndentedJSON(http.StatusOK, updatestudent)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "student not found"})
}
func delete_student(c *gin.Context) {
	name := c.Param("name")
	if name == " " {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name not found"})
		return
	}
	for i, detials := range student {
		if detials.Name == name {
			student = append(student[:i], student[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "student deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "student name id not found"})
}
