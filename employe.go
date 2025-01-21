package main

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtkey = []byte("my_secret_key")

//	type credentials struct {
//		Email    string `json:"email"`
//		Password string `json:"password"`
//	}
var users = map[string]string{
	"venkatesh999@gmail.com": "Venkey83$",
}

type Cliams struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func main() {
	g := gin.Default()
	g.Use(gin.Logger())
	g.POST("/login", loginhandler)
	g.POST("/logout", logouthandler)
	g.POST("/forgot-password", forgotpasswordhandller)
	g.POST("/reset-password", resethandler)
	g.Run(":1010")

}
func loginhandler(c *gin.Context) {
	var request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	storedpassword, exists := users[request.Email]
	if !exists || storedpassword != request.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}
	expertion_time := time.Now().Add(15 * time.Minute)
	cliams := &Cliams{
		Email: request.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expertion_time.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cliams)
	tokenstring, err := token.SignedString(jwtkey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create token"})
		return
	}
	c.JSON(http.StatusOK, tokenstring)
}
func logouthandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "logout sucessfully"})
}
func forgotpasswordhandller(c *gin.Context) {
	var request struct {
		Email string `json:"email"`
	}
	//var request credentials
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "invalid request"})
		return
	}
	if _, exists := users[request.Email]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "email not found"})
		return
	}
	//fmt.Printf("password reset link set %s\n", request.Email)
	c.JSON(http.StatusOK, gin.H{"message": "password reset link sent successfully"})

}
func resethandler(c *gin.Context) {
	var request struct {
		Email       string `json:"email"`
		Newpassword string `json:"new_password"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	if _, exists := users[request.Email]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "email not found"})
		return
	}
	users[request.Email] = request.Newpassword
	//fmt.Printf("reset password for employe %s\n", request.Email)
	c.JSON(http.StatusOK, gin.H{"message": "password reset successfully"})
}
