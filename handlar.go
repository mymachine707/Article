package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func remove(slice []Article, s int) []Article {
	return append(slice[:s], slice[s+1:]...)
}

// InMemoryArticleData - data base article
var InMemoryArticleData []Article

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
				"message": "GetArticleById",
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

// GetArticleList godoc
// @Summary      List articles
// @Description  GetArticleList
// @Tags         articless
// @Accept       json
// @Produce      json
// @Success      200  {object}  JSONResult{data=[]Article}
// @Router       /v2/article/{id} [get]
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
