package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = []User{
	{ID: 1, Name: "Sai", Email: "sai@example.com"},
	{ID: 2, Name: "Ganesh", Email: "ganesh@example.com"},
}
var nextID = 3

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Gin CRUD is working!"})
	})

	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, users)
	})

	r.POST("/users", func(c *gin.Context) {
		var newUser User
		if err := c.BindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}
		newUser.ID = nextID
		nextID++
		users = append(users, newUser)
		c.JSON(http.StatusCreated, newUser)
	})

	r.PUT("/users/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		var updatedUser User
		if err := c.BindJSON(&updatedUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		for i, u := range users {
			if u.ID == id {
				users[i].Name = updatedUser.Name
				users[i].Email = updatedUser.Email
				c.JSON(http.StatusOK, users[i])
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	})

	r.DELETE("/users/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		for i, u := range users {
			if u.ID == id {
				users = append(users[:i], users[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	})

	r.Run(":8080")
}
