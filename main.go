package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Person ...
type Person struct {
	Firstname, Lastname string // Compact by combining the various fields of the same type
}

// Content ...
type Content struct {
	Title string
	Body  string
}

// Article ...
type Article struct {
	ID        int
	Content          // Promoted fields
	Author    Person // Nested structs
	CreatedAt *time.Time
}

func main() {

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Create
	r.POST("/article", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Article Create ",
		})
	})

	// GetArticleById
	r.GET("/article/:id", func(c *gin.Context) {
		idStr := c.Param("id")

		fmt.Println(idStr)

		c.JSON(http.StatusOK, gin.H{
			"message": "GetArticleById",
		})
	})

	// GetList
	r.GET("/article", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Article GetList",
		})
	})

	// Update
	r.PUT("/article", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Article Update",
		})
	})

	// Delete
	r.DELETE("/article/:id", func(c *gin.Context) {
		idStr := c.Param("id")

		fmt.Println(idStr)

		c.JSON(http.StatusOK, gin.H{
			"message": "GetArticleById",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
