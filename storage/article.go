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
		if v.ID == id {
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
	return result, errors.New("Article not found!")
}

// GetArticleList ...
func GetArticleList() (resp []models.Article, err error) {
	resp = InMemoryArticleData
	return resp, err
}

// UpdateArticle ...
func UpdateArticle(article models.Article) error {
	for i, v := range InMemoryArticleData {
		if v.ID == article.ID {
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
			InMemoryArticleData = remove(InMemoryArticleData, i)
			return nil
		}
	}
	return errors.New("Cannot delete article becouse Article not found")
}

func remove(slice []models.Article, s int) []models.Article {
	return append(slice[:s], slice[s+1:]...)
}
