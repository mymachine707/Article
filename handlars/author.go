package handlars

import (
	"mymachine707/models"
	"mymachine707/storage"
	"net/http"

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
// @Success     201     {object} models.JSONResult{data=models.Author}
// @Failure     400     {object} models.JSONErrorResponse
// @Router      /v2/author [post]
func CreatAuthor(c *gin.Context) {

	var body models.CreateAuthorModul

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	// validation should be here

	// create new author
	id := uuid.New()
	err := storage.AddAuthor(id.String(), body)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	author, err := storage.GetAuthorByID(id.String()) // maqsad tekshirish rostan  ham create bo'ldimi?

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
func GetAuthorByID(c *gin.Context) {

	idStr := c.Param("id")

	// validation

	author, err := storage.GetAuthorByID(idStr)

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
// @Success     200 {object} models.JSONResult{data=[]models.Author}
// @Router      /v2/author/ [get]
func GetAuthorList(c *gin.Context) {

	authorList, err := storage.GetAuthorList()

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
// @Param       author body     models.CreateAuthorModul true "Author body"
// @Success     201     {object} models.JSONResult{data=[]models.Author}
// @Failure     400     {object} models.JSONErrorResponse
// @Router      /v2/author/ [put]
func AuthorUpdate(c *gin.Context) {
	var author models.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	// my work change code ... mst

	err := storage.UpdateAuthor(author)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Author Update",
		"data":    storage.InMemoryAuthorData,
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
func DeleteAuthor(c *gin.Context) {
	idStr := c.Param("id")

	// my code change ...
	err := storage.DeleteAuthor(idStr)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Author Deleted",
	})

}
