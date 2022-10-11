package storage

import "mymachine707/models"

// Interfaces ...
type Interfaces interface {
	// For Article
	AddArticle(id string, entity models.CreateArticleModul) error
	GetArticleByID(id string) (models.PackedArticleModel, error)
	GetArticleList(offset, limit int, search string) (resp []models.Article, err error)
	UpdateArticle(article models.UpdateArticleModul) error
	DeleteArticle(idStr string) error

	// For Author
	AddAuthor(id string, entity models.CreateAuthorModul) error
	GetAuthorByID(id string) (models.Author, error)
	GetAuthorList(offset, limit int, serach string) (resp []models.Author, err error)
	UpdateAuthor(author models.UpdateAuthorModul) error
	DeleteAuthor(idStr string) error
}
