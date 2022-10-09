package storage

import (
	"errors"
	"mymachine707/models"
	"time"
)

// InMemoryArticleData - data base article
var InMemoryArticleData []models.Article

// AddArticle ...
func AddArticle(id string, entity models.CreateArticleModul) error {

	var article models.Article

	article.ID = id
	article.Content = entity.Content
	article.AuthorID = entity.AuthorID
	article.CreatedAt = time.Now()

	InMemoryArticleData = append(InMemoryArticleData, article)

	return nil
}

// GetArticleByID ...
func GetArticleByID(id string) (models.PackedArticleModel, error) {
	var result models.PackedArticleModel

	for _, v := range InMemoryArticleData {
		if v.ID == id && v.DeletedAt != nil {
			return result, errors.New("article already deleted")
		}
		if v.ID == id && v.DeletedAt == nil {
			author, err := GetAuthorByID(v.AuthorID)
			if err != nil {
				return result, err
			}
			result.ID = v.ID
			result.Author = author
			result.Content = v.Content
			result.CreatedAt = v.CreatedAt
			result.UpdatedAt = v.UpdatedAt
			result.DeletedAt = v.DeletedAt
			return result, nil
		}
	}
	return result, errors.New("article not found")
}

// GetArticleList ...
func GetArticleList() (resp []models.Article, err error) {
	for _, v := range InMemoryArticleData {
		if v.DeletedAt == nil {
			resp = append(resp, v)
		}
	}

	return resp, err
}

// UpdateArticle ...
func UpdateArticle(article models.Article) error {
	for i, v := range InMemoryArticleData {
		if v.ID == article.ID && v.DeletedAt == nil {
			article.CreatedAt = v.CreatedAt
			t := time.Now()
			article.UpdatedAt = &t
			InMemoryArticleData[i] = article
			return nil
		}
	}
	return errors.New("Cannot Update article")
}

// DeleteArticle ...
func DeleteArticle(idStr string) error {

	for i, v := range InMemoryArticleData {
		if v.ID == idStr {
			if v.DeletedAt != nil {
				return errors.New("article already deleted")
			}
			// bu kod article hard delete qilish uchun :
			// InMemoryArticleData = remove(InMemoryArticleData, i)

			// bu kod soft delete uchun:
			t := time.Now()
			v.DeletedAt = &t
			InMemoryArticleData[i] = v
			return nil
		}
	}
	return errors.New("Cannot delete article becouse Article not found")
}

func remove(slice []models.Article, s int) []models.Article {
	return append(slice[:s], slice[s+1:]...)
}
