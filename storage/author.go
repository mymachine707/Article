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
func UpdateAuthor(author models.Author) error {
	for i, v := range InMemoryAuthorData {
		if v.ID == author.ID {
			author.CreatedAt = v.CreatedAt
			t := time.Now()
			author.UpdatedAt = &t
			InMemoryAuthorData[i] = author
			return nil
		}
	}
	return errors.New("Cannot Update author")
}

// DeleteAuthor ...
func DeleteAuthor(idStr string) error {

	for i, v := range InMemoryAuthorData {
		if v.ID == idStr {
			InMemoryAuthorData = removeAuthorDelete(InMemoryAuthorData, i)
			return nil
		}
	}
	return errors.New("Cannot delete article becouse Author not found")
}

func removeAuthorDelete(slice []models.Author, s int) []models.Author {
	return append(slice[:s], slice[s+1:]...)
}
