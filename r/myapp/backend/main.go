package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

type User struct {
	Username, Password string
}

var DB *gorm.DB

func Conn() {
	connDb := "root:root@/react"
	Db, err := gorm.Open(mysql.Open(connDb), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = Db

	err = DB.AutoMigrate(&User{})
	if err != nil {
		return
	}

}

func Register(c *gin.Context) {
	var user User

	// Bind JSON data from the request
	if err := c.BindJSON(&user); err != nil {
		// Handle validation errors
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create the user record in the database
	if err := DB.Create(&user).Error; err != nil {
		// Handle database errors
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Respond with a success message
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func main() {

	Conn()

	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*") // Allow requests from any origin
		c.Next()
	})

	router.LoadHTMLGlob("react-redirect/public/index.html")

	//router.LoadHTMLGlob("C:\\Users\\sri.avanigadda\\Desktop\\r\\myapp\\react-redirect\\src/*.html")

	router.POST("/create", Register)

	err := router.Run()

	if err != nil {
		return
	}
}
