package handlars

import (
	"mymachine707/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func remove(slice []models.Article, s int) []models.Article {
	return append(slice[:s], slice[s+1:]...)
}

// InMemoryArticleData - data base article
var InMemoryArticleData []models.Article

// CreatArticle godoc
// @Summary     Creat Article
// @Description Creat a new article
// @Tags        article
// @Accept      json
// @Produce     json
// @Param       article body     models.CreateArticleModul true "Article body"
// @Success     201     {object} models.Article
// @Failure     400     {object} models.JSONErrorResponse
// @Router      /v2/article [post]
func CreatArticle(c *gin.Context) {

	var article models.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	// validation should be here

	id := uuid.New()
	article.ID = id.String()
	article.CreatedAt = time.Now()

	InMemoryArticleData = append(InMemoryArticleData, article)

	c.JSON(http.StatusCreated, models.JSONResult{
		Data:    InMemoryArticleData,
		Message: "CreatArticle",
	})
}

// GetArticleByID godoc
// @Summary     GetArticleByID
// @Description get an article by id
// @Tags        article
// @Accept      json
// @Produce     json
// @Param       id  path     string true "Article id"
// @Success     201 {object} models.JSONResult{data=models.Article}
// @Failure     400 {object} models.JSONErrorResponse
// @Router      /v2/article/{id} [get]
func GetArticleByID(c *gin.Context) {
	idStr := c.Param("id")

	for _, v := range InMemoryArticleData {
		if v.ID == idStr {
			c.JSON(http.StatusOK, models.JSONResult{
				Data:    v,
				Message: "GetArticleByID",
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, models.JSONErrorResponse{
		Error: "GetArticleById || NOT FOUND",
	})

}

// GetArticleList godoc
// @Summary     List articles
// @Description GetArticleList
// @Tags        article
// @Accept      json
// @Produce     json
// @Success     200 {object} models.JSONResult{data=[]models.Article}
// @Router      /v2/article/ [get]
func GetArticleList(c *gin.Context) {
	c.JSON(http.StatusOK, models.JSONResult{
		Data:    InMemoryArticleData,
		Message: "Article GetList",
	})
}

// ArticleUpdate godoc
// @Summary     My work !!! -- Update Article
// @Description Update Article
// @Tags        article
// @Accept      json
// @Produce     json
// @Param       article body     models.CreateArticleModul true "Article body"
// @Success     201     {object} models.JSONResult{data=[]models.Article}
// @Failure     400     {object} models.JSONErrorResponse
// @Router      /v2/article/ [put]
func ArticleUpdate(c *gin.Context) {
	var article models.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
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

	c.JSON(http.StatusNotFound, models.JSONErrorResponse{
		Error: "Update || NOT FOUND",
	})

}

// DeleteArticle godoc
// @Summary     My work!!! -- Delete Article
// @Description get element by id and delete this article
// @Tags        article
// @Accept      json
// @Produce     json
// @Param       id  path     string true "Article id"
// @Success     201 {object} models.JSONResult{data=models.Article}
// @Failure     400 {object} models.JSONErrorResponse
// @Router      /v2/article/{id} [delete]
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

	c.JSON(http.StatusNotFound, models.JSONErrorResponse{
		Error: "Delete element || NOT FOUND",
	})
}
