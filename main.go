package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Person ...
type Person struct {
	Firstname string `json:"firstname" ` // binding:"required"
	Lastname  string `json:"firstname" ` // binding:"required"
}

// Content ...
type Content struct {
	Title string `json:"title" ` // binding:"required"
	Body  string `json:"body" `  // binding:"required"
}

// Article ...
type Article struct {
	ID        int        `json:"id"`
	Content              // Promoted fields
	Author    Person     `json:"a" ` // binding:"required"
	CreatedAt *time.Time `json:"created_at"`
}

func main() {

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	var InMemoryArticleData []Article

	// Create
	r.POST("/article", func(c *gin.Context) {

		var article Article
		if err := c.ShouldBindJSON(&article); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		t := time.Now()
		article.CreatedAt = &t
		InMemoryArticleData = append(InMemoryArticleData, article)

		c.JSON(http.StatusOK, gin.H{
			"data":    InMemoryArticleData,
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
