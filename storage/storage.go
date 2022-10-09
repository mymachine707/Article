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
	article.Author = entity.Author
	article.CreatedAt = time.Now()

	InMemoryArticleData = append(InMemoryArticleData, article)

	return nil
}

// GetArticleByID ...
func GetArticleByID(id string) (models.Article, error) {
	for _, v := range InMemoryArticleData {
		if v.ID == id {
			return v, nil
		}
	}
	return models.Article{}, errors.New("Article not found!")
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
