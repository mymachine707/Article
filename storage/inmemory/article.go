package inmemory

import (
	"errors"
	"mymachine707/models"
	"strings"
	"time"
)

// AddArticle ...
func (IM InMemory) AddArticle(id string, entity models.CreateArticleModul) error {
	var article models.Article

	if id == "" {
		return errors.New("id must exist")
	}

	article.ID = id
	article.Content = entity.Content

	author, err := IM.GetAuthorByID(entity.AuthorID)

	if err != nil {
		return err
	}

	article.AuthorID = author.ID
	article.CreatedAt = time.Now()

	IM.Db.InMemoryArticleData = append(IM.Db.InMemoryArticleData, article)

	return nil
}

// GetArticleByID ...
func (IM InMemory) GetArticleByID(id string) (models.PackedArticleModel, error) {
	var result models.PackedArticleModel

	for _, v := range IM.Db.InMemoryArticleData {
		if v.ID == id && v.DeletedAt != nil {
			return result, errors.New("article already deleted")
		}
		if v.ID == id && v.DeletedAt == nil {
			author, err := IM.GetAuthorByID(v.AuthorID)
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
func (IM InMemory) GetArticleList(offset, limit int, search string) (resp []models.Article, err error) {
	off := 0
	c := 0

	for _, v := range IM.Db.InMemoryArticleData {
		if v.DeletedAt == nil && (strings.Contains(v.Title, search) || strings.Contains(v.Body, search)) {

			if offset <= off {
				c++
				resp = append(resp, v)
			}
			if limit <= c {
				break
			}
			off++
		}
	}

	return resp, err
}

// UpdateArticle ...
func (IM InMemory) UpdateArticle(article models.UpdateArticleModul) error {

	for i, v := range IM.Db.InMemoryArticleData {
		if v.ID == article.ID && v.DeletedAt == nil {

			v.Content = article.Content
			t := time.Now()
			v.UpdatedAt = &t

			IM.Db.InMemoryArticleData[i] = v

			return nil
		}
	}
	return errors.New("article not found")
}

// DeleteArticle ...
func (IM InMemory) DeleteArticle(idStr string) error {

	for i, v := range IM.Db.InMemoryArticleData {
		if v.ID == idStr {
			if v.DeletedAt != nil {
				return errors.New("article already deleted")
			}
			// bu kod article hard delete qilish uchun :
			// IM.Db.InMemoryArticleData = remove(IM.Db.InMemoryArticleData, i)

			// bu kod soft delete uchun:
			t := time.Now()
			v.DeletedAt = &t
			IM.Db.InMemoryArticleData[i] = v
			return nil
		}
	}
	return errors.New("Cannot delete article becouse Article not found")
}

func (IM InMemory) remove(slice []models.Article, s int) []models.Article {
	return append(slice[:s], slice[s+1:]...)
}
