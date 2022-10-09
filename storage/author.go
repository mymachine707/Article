package storage

import (
	"errors"
	"mymachine707/models"
	"time"
)

// InMemoryAuthorData - data base author
var InMemoryAuthorData []models.Author

// AddAuthor ...
func AddAuthor(id string, entity models.CreateAuthorModul) error {

	var author models.Author

	author.ID = id
	author.Firstname = entity.Firstname
	author.Lastname = entity.Lastname
	author.CreatedAt = time.Now()

	InMemoryAuthorData = append(InMemoryAuthorData, author)

	return nil
}

// GetAuthorByID ...
func GetAuthorByID(id string) (models.Author, error) {
	var result models.Author

	for _, v := range InMemoryAuthorData {
		if v.ID == id {
			result = v
			return result, nil
		}
	}
	return result, errors.New("author not found")
}

// GetAuthorList ...
func GetAuthorList() (resp []models.Author, err error) {
	resp = InMemoryAuthorData
	return resp, err
}

// UpdateAuthor ...
func UpdateAuthor(article models.Author) error {
	for i, v := range InMemoryAuthorData {
		if v.ID == article.ID {
			article.CreatedAt = v.CreatedAt
			t := time.Now()
			article.UpdatedAt = &t
			InMemoryAuthorData[i] = article
			return nil
		}
	}
	return errors.New("Cannot Update article")
}

// DeleteAuthor ...
func DeleteAuthor(idStr string) error {

	for i, v := range InMemoryArticleData {
		if v.ID == idStr {
			InMemoryArticleData = remove(InMemoryArticleData, i)
			return nil
		}
	}
	return errors.New("Cannot delete article becouse Article not found")
}
