package handlars

import (
	"mymachine707/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreatAuthor godoc
// @Summary     Creat Author
// @Description Creat a new author
// @Tags        author
// @Accept      json
// @Produce     json
// @Param       author body     models.CreateAuthorModul true "Author body"
// @Success     201    {object} models.JSONResult{data=models.Author}
// @Failure     400    {object} models.JSONErrorResponse
// @Router      /v2/author [post]
func (h *Handler) CreatAuthor(c *gin.Context) {

	var body models.CreateAuthorModul

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	// validation should be here

	// create new author

	id := uuid.New()
	err := h.Stg.AddAuthor(id.String(), body)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}
	author, err := h.Stg.GetAuthorByID(id.String()) // maqsad tekshirish rostan  ham create bo'ldimi?
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, models.JSONResult{
		Message: "CreatAuthor",
		Data:    author,
	})
}

// GetAuthorByID godoc
// @Summary     GetAuthorByID
// @Description get an author by id
// @Tags        author
// @Accept      json
// @Produce     json
// @Param       id  path     string true "Author id"
// @Success     201 {object} models.JSONResult{data=models.PackedAuthorModel}
// @Failure     400 {object} models.JSONErrorResponse
// @Router      /v2/author/{id} [get]
func (h *Handler) GetAuthorByID(c *gin.Context) {

	idStr := c.Param("id")

	// validation

	author, err := h.Stg.GetAuthorByID(idStr)

	if err != nil {
		c.JSON(http.StatusNotFound, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResult{
		Message: "OK",
		Data:    author,
	})
}

// GetAuthorList godoc
// @Summary     List authors
// @Description GetAuthorList
// @Tags        author
// @Accept      json
// @Produce     json
// @Param       offset query    int    false "0"
// @Param       limit  query    int    false "100"
// @Param       search query    string false "search exapmle"
// @Success     200    {object} models.JSONResult{data=[]models.Author}
// @Router      /v2/author/ [get]
func (h *Handler) GetAuthorList(c *gin.Context) {

	offsetStr := c.DefaultQuery("offset", "0")
	limitStr := c.DefaultQuery("limit", "100")
	search := c.DefaultQuery("search", "")

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	authorList, err := h.Stg.GetAuthorList(offset, limit, search)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResult{
		Data:    authorList,
		Message: "GetList OK",
	})
}

// AuthorUpdate godoc
// @Summary     My work !!! -- Update Author
// @Description Update Author
// @Tags        author
// @Accept      json
// @Produce     json
// @Param       author body     models.UpdateAuthorModul true "Author body"
// @Success     201    {object} models.JSONResult{data=[]models.Author}
// @Failure     400    {object} models.JSONErrorResponse
// @Router      /v2/author/ [put]
func (h *Handler) AuthorUpdate(c *gin.Context) {
	var body models.UpdateAuthorModul
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	// my work change code ... mst

	err := h.Stg.UpdateAuthor(body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	res, err := h.Stg.GetAuthorByID(body.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Author Update",
		"data":    res,
	})

}

// DeleteAuthor godoc
// @Summary     My work!!! -- Delete Author
// @Description get element by id and delete this author
// @Tags        author
// @Accept      json
// @Produce     json
// @Param       id  path     string true "Author id"
// @Success     201 {object} models.JSONResult{data=models.Author}
// @Failure     400 {object} models.JSONErrorResponse
// @Router      /v2/author/{id} [delete]
func (h *Handler) DeleteAuthor(c *gin.Context) {
	idStr := c.Param("id")

	author, err := h.Stg.GetAuthorByID(idStr)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	// my code change ...
	err = h.Stg.DeleteAuthor(idStr)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Author Deleted",
		"data":    author,
	})

}
