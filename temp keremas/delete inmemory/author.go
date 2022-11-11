package inmemory

import (
	"errors"
	"mymachine707/models"
	"strings"
	"time"
)

// AddAuthor ...
func (IM InMemory) AddAuthor(id string, entity models.CreateAuthorModul) error {
	var author models.Author

	if id == "" {
		return errors.New("id must exist")
	}
	if entity.Firstname == "" {
		return errors.New("Firstname must exist")
	}
	if entity.Lastname == "" {
		return errors.New("Lastname must exist")
	}

	author.ID = id
	author.Firstname = entity.Firstname
	author.Lastname = entity.Lastname
	author.Middlename = entity.Middlename
	author.CreatedAt = time.Now()

	IM.Db.InMemoryAuthorData = append(IM.Db.InMemoryAuthorData, author)

	return nil
}

// GetAuthorByID ...
func (IM InMemory) GetAuthorByID(id string) (models.Author, error) {
	var result models.Author

	for _, v := range IM.Db.InMemoryAuthorData {
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
func (IM InMemory) GetAuthorList(offset, limit int, serach string) (resp []models.Author, err error) {
	off := 0
	c := 0

	for _, v := range IM.Db.InMemoryAuthorData {
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
func (IM InMemory) UpdateAuthor(author models.UpdateAuthorModul) error {

	for i, v := range IM.Db.InMemoryAuthorData {
		if v.ID == author.ID && v.DeletedAt == nil {

			v.Firstname = author.Firstname
			v.Lastname = author.Lastname
			t := time.Now()
			v.UpdatedAt = &t

			IM.Db.InMemoryAuthorData[i] = v

			return nil
		}
	}
	return errors.New("author not found")
}

// DeleteAuthor ...
func (IM InMemory) DeleteAuthor(idStr string) error {

	for i, v := range IM.Db.InMemoryAuthorData {
		if v.ID == idStr {
			// hard delete uchun kod
			//IM.Db.InMemoryAuthorData = removeAuthorDelete(IM.Db.InMemoryAuthorData, i)

			// soft delete uchun kod
			if v.DeletedAt != nil {
				return errors.New("author already deleted")
			}
			t := time.Now()
			v.DeletedAt = &t
			IM.Db.InMemoryAuthorData[i] = v
			return nil
		}
	}
	return errors.New("Cannot delete article becouse Author not found")
}

// hard delete uchun kod
// func (IM InMemory) removeAuthorDelete(slice []models.Author, s int) []models.Author {
// 	return append(slice[:s], slice[s+1:]...)
// }
