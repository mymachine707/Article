package handlars

import (
	"mymachine707/models"
	"net/http"
	"strconv"

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
func (h *Handler) CreatArticle(c *gin.Context) {

	var body models.CreateArticleModul

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	// validation should be here

	// create new article
	id := uuid.New()
	err := h.Stg.AddArticle(id.String(), body)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	article, err := h.Stg.GetArticleByID(id.String()) // maqsad tekshirish rostan  ham create bo'ldimi?

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
func (h *Handler) GetArticleByID(c *gin.Context) {

	idStr := c.Param("id")

	// validation

	article, err := h.Stg.GetArticleByID(idStr)

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
// @Param       offset query    int    false "0"   default(A)
// @Param       limit  query    int    false "100" default(A)
// @Param       search query    string false "s"   default(A)
// @Success     200    {object} models.JSONResult{data=[]models.Article}
// @Router      /v2/article/ [get]
func (h *Handler) GetArticleList(c *gin.Context) {

	offsetStr := c.DefaultQuery("offset", "0")
	limitStr := c.DefaultQuery("limit", "100")

	searchStr := c.DefaultQuery("search", "")

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	articleList, err := h.Stg.GetArticleList(offset, limit, searchStr)

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
// @Param       article body     models.UpdateArticleModul true "Article body"
// @Success     201     {object} models.JSONResult{data=[]models.Article}
// @Failure     400     {object} models.JSONErrorResponse
// @Router      /v2/article/ [put]
func (h *Handler) ArticleUpdate(c *gin.Context) {
	var body models.UpdateArticleModul
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	// my work change code ... mst

	err := h.Stg.UpdateArticle(body)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{ //!
			Error: err.Error(),
		})
		return
	}

	res, err := h.Stg.GetArticleByID(body.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Article Update",
		"data":    res,
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
func (h *Handler) DeleteArticle(c *gin.Context) {
	idStr := c.Param("id")

	article, err := h.Stg.GetArticleByID(idStr)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	// my code change ...
	err = h.Stg.DeleteArticle(idStr)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Article Deleted",
		"data":    article,
	})
}
