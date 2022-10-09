package storage

import (
	"errors"
	"mymachine707/models"
	"time"
)

// InMemoryAuthorData - data base author
var InMemoryAuthorData []models.Author

// AddAuthor ...
func AddAuthor(id string, entity models.AuthorCreateModel) error {

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
	return result, errors.New("Author not found!")
}
