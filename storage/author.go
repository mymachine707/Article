package storage

import (
	"errors"
	"mymachine707/models"
	"strings"
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
		if v.ID == id && v.DeletedAt != nil {
			return result, errors.New("author already deleted")
		}
		if v.ID == id && v.DeletedAt == nil {
			result = v
			return result, nil
		}
	}
	return result, errors.New("author not found")
}

// GetAuthorList ...
func GetAuthorList(offset, limit int, serach string) (resp []models.Author, err error) {
	off := 0
	c := 0

	for _, v := range InMemoryAuthorData {
		if v.DeletedAt == nil && (strings.Contains(v.Firstname, serach) || strings.Contains(v.Lastname, serach)) {
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

// UpdateAuthor ...
func UpdateAuthor(author models.UpdateAuthorModul) error {

	for i, v := range InMemoryAuthorData {
		if v.ID == author.ID && v.DeletedAt == nil {

			v.Firstname = author.Firstname
			v.Lastname = author.Lastname
			t := time.Now()
			v.UpdatedAt = &t

			InMemoryAuthorData[i] = v

			return nil
		}
	}
	return errors.New("author not found")
}

// DeleteAuthor ...
func DeleteAuthor(idStr string) error {

	for i, v := range InMemoryAuthorData {
		if v.ID == idStr {
			// hard delete uchun kod
			//InMemoryAuthorData = removeAuthorDelete(InMemoryAuthorData, i)

			// soft delete uchun kod
			if v.DeletedAt != nil {
				return errors.New("author already deleted")
			}
			t := time.Now()
			v.DeletedAt = &t
			InMemoryAuthorData[i] = v
			return nil
		}
	}
	return errors.New("Cannot delete article becouse Author not found")
}

func removeAuthorDelete(slice []models.Author, s int) []models.Author {
	return append(slice[:s], slice[s+1:]...)
}
