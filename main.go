package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Person ...
type Person struct {
	Firstname string `json:"firstname" binding:"required" `
	Lastname  string `json:"lastname" binding:"required"`
}

// Content ...
type Content struct {
	Title string `json:"title" binding:"required" `
	Body  string `json:"body" binding:"required"`
}

// Article ...
type Article struct {
	ID        string     `json:"id"`
	Content              // Promoted fields
	Author    Person     `json:"author" binding:"required" `
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

// InMemoryArticleData - data base article
var InMemoryArticleData []Article

func main() {
	InMemoryArticleData = make([]Article, 0)
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Create
	r.POST("/article", func(c *gin.Context) {

		var article Article
		if err := c.ShouldBindJSON(&article); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id := uuid.New()

		article.ID = id.String()

		article.CreatedAt = time.Now()
		InMemoryArticleData = append(InMemoryArticleData, article)

		c.JSON(http.StatusOK, gin.H{
			"data":    InMemoryArticleData,
			"message": "Article Create ",
		})
	})

	// GetArticleById
	r.GET("/article/:id", func(c *gin.Context) {
		idStr := c.Param("id")

		for _, v := range InMemoryArticleData {
			if v.ID == idStr {
				c.JSON(http.StatusOK, gin.H{
					"message": "GetArticleById ",
					"data":    v,
				})
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{
			"message": "GetArticleById || NOT FOUND",
			"data":    nil,
		})

	})

	// GetList
	r.GET("/article", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Article GetList",
			"data":    InMemoryArticleData,
		})
	})

	// Update
	r.PUT("/article", func(c *gin.Context) {
		var article Article
		if err := c.ShouldBindJSON(&article); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		for i, v := range InMemoryArticleData {
			if v.ID == article.ID {
				t := time.Now()
				article.UpdatedAt = &t
				InMemoryArticleData[i] = article
				c.JSON(http.StatusNotFound, gin.H{
					"message": "GetArticleById || NOT FOUND",
					"data":    InMemoryArticleData,
				})
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{
			"message": "Update || NOT FOUND",
			"data":    nil,
		})

	})

	// Delete
	r.DELETE("/article/:id", func(c *gin.Context) {
		idStr := c.Param("id")

		for i, v := range InMemoryArticleData {
			if v.ID == idStr {
				c.JSON(http.StatusNotFound, gin.H{
					"message": "GetArticleById || NOT FOUND",
					"data":    v,
				})
				InMemoryArticleData = remove(InMemoryArticleData, i)
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Article || Delete || NOT FOUND",
			"data":    nil,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func remove(slice []Article, s int) []Article {
	return append(slice[:s], slice[s+1:]...)
}
