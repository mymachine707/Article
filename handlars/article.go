package handlars

import (
	"mymachine707/models"
	"mymachine707/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreatArticle godoc
// @Summary     Creat Article
// @Description Creat a new article
// @Tags        article
// @Accept      json
// @Produce     json
// @Param       article body     models.CreateArticleModul true "Article body"
// @Success     201     {object} models.JSONResult{data=models.Article}
// @Failure     400     {object} models.JSONErrorResponse
// @Router      /v2/article [post]
func CreatArticle(c *gin.Context) {

	var body models.CreateArticleModul

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	// validation should be here

	// create new article
	id := uuid.New()
	err := storage.AddArticle(id.String(), body)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	article, err := storage.GetArticleByID(id.String()) // maqsad tekshirish rostan  ham create bo'ldimi?

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.JSONResult{
		Message: "CreatArticle",
		Data:    article,
	})
}

// GetArticleByID godoc
// @Summary     GetArticleByID
// @Description get an article by id
// @Tags        article
// @Accept      json
// @Produce     json
// @Param       id  path     string true "Article id"
// @Success     201 {object} models.JSONResult{data=models.PackedArticleModel}
// @Failure     400 {object} models.JSONErrorResponse
// @Router      /v2/article/{id} [get]
func GetArticleByID(c *gin.Context) {

	idStr := c.Param("id")

	// validation

	article, err := storage.GetArticleByID(idStr)

	if err != nil {
		c.JSON(http.StatusNotFound, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResult{
		Message: "OK",
		Data:    article,
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

	articleList, err := storage.GetArticleList()

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResult{
		Data:    articleList,
		Message: "GetList OK",
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

	// my work change code ... mst

	err := storage.UpdateArticle(article)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Article Update",
		"data":    storage.InMemoryArticleData,
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

	// my code change ...
	err := storage.DeleteArticle(idStr)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Article Deleted",
	})

}
