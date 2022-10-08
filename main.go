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
	r.POST("/article", CreatArticle)

	// GetArticleById
	r.GET("/article/:id", GetArticleByID)

	// GetList
	r.GET("/article", GetArticleList)

	// Update
	r.PUT("/article", ArticleUpdate)

	// Delete
	r.DELETE("/article/:id", DeleteArticle)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func remove(slice []Article, s int) []Article {
	return append(slice[:s], slice[s+1:]...)
}

//CreatArticle ...
func CreatArticle(c *gin.Context) {

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
}

// GetArticleByID ...
func GetArticleByID(c *gin.Context) {
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

}

// GetArticleList ...
func GetArticleList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Article GetList",
		"data":    InMemoryArticleData,
	})
}

// ArticleUpdate ...
func ArticleUpdate(c *gin.Context) {
	var article Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, v := range InMemoryArticleData {
		if v.ID == article.ID {
			article.CreatedAt = v.CreatedAt
			t := time.Now()
			article.UpdatedAt = &t
			InMemoryArticleData[i] = article
			c.JSON(http.StatusOK, gin.H{
				"message": "Article Update",
				"data":    InMemoryArticleData,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "Update || NOT FOUND",
		"data":    nil,
	})

}

// DeleteArticle ...
func DeleteArticle(c *gin.Context) {
	idStr := c.Param("id")

	for i, v := range InMemoryArticleData {
		if v.ID == idStr {
			c.JSON(http.StatusOK, gin.H{
				"message": "Article Deleted",
				"data":    v,
			})
			InMemoryArticleData = remove(InMemoryArticleData, i)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "Article || Delete || NOT FOUND",
		"data":    nil,
	})
}
