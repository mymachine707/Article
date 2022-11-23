package postgres

import (
	"errors"
	"fmt"
	"mymachine707/models"
)

var err error

// AddAuthor ...
func (stg Postgres) AddAuthor(id string, entity models.CreateAuthorModul) error {
	if id == "" {
		return errors.New("id must exist")
	}
	fname := entity.Firstname + " " + entity.Lastname + " " + entity.Middlename

	_, err = stg.db.Exec(`INSERT INTO author (
		id,
		firstname,
		lastname,
		middlename,
		fullname
		) VALUES(
		$1,
		$2,
		$3,
		$4,
		$5
	)`,
		id,
		entity.Firstname,
		entity.Lastname,
		entity.Middlename,
		fname,
	)

	if err != nil {
		return err
	}

	return nil
}

// GetAuthorByID ...
func (stg Postgres) GetAuthorByID(id string) (models.Author, error) {
	var a models.Author

	if id == "" {
		return a, errors.New("id must exist")
	}

	err := stg.db.QueryRow(`SELECT
		au.id,
		au.firstname,
		au.lastname,
		au.middlename,
		au.fullname,
		au.created_at,
		au.updated_at,
		au.deleted_at
	FROM author AS au WHERE id=$1 AND deleted_at is null`, id).Scan(
		&a.ID,
		&a.Firstname,
		&a.Lastname,
		&a.Middlename,
		&a.Fullname,
		&a.CreatedAt,
		&a.UpdatedAt,
		&a.DeletedAt,
	)

	if err != nil {
		return a, err
	}

	return a, nil
}

// GetAuthorList ...
func (stg Postgres) GetAuthorList(offset, limit int, search string) (resp []models.Author, err error) {

	rows, err := stg.db.Queryx(`
	
	Select * from author WHERE 

		((firstname ILIKE '%' || $1 || '%') OR (lastname ILIKE '%' || $1 || '%') OR 
		(middlename ILIKE '%' || $1 || '%') OR 
		(fullname ILIKE '%' || $1 || '%'))
		AND deleted_at is null 
		LIMIT $2 
		OFFSET $3`, search, limit, offset)

	if err != nil {
		return resp, err
	}

	for rows.Next() {
		var a models.Author

		err = rows.Scan(
			&a.ID,
			&a.Firstname,
			&a.Lastname,
			&a.CreatedAt,
			&a.UpdatedAt,
			&a.DeletedAt,
			&a.Middlename,
			&a.Fullname,
		)
		fmt.Println(a.UpdatedAt)
		if err != nil {
			return resp, err
		}

		resp = append(resp, a)

	}

	return resp, nil
}

// UpdateAuthor ...
func (stg Postgres) UpdateAuthor(author models.UpdateAuthorModul) error {

	fname := author.Firstname + " " + author.Lastname + " " + author.Middlename

	rows, err := stg.db.NamedExec(`Update author set firstname=:f, lastname=:l, middlename=:m, fullname=:fn,updated_at=now() Where id=:id  and deleted_at is null`, map[string]interface{}{
		"id": author.ID,
		"f":  author.Firstname,
		"l":  author.Lastname,
		"m":  author.Middlename,
		"fn": fname,
	})

	if err != nil {
		return err
	}

	n, err := rows.RowsAffected()

	if err != nil {
		return err
	}

	if n > 0 {
		return nil
	}

	return errors.New("author not found")
}

// DeleteAuthor ...
func (stg Postgres) DeleteAuthor(idStr string) error {

	rows, err := stg.db.Exec(`UPDATE author SET deleted_at=now() Where id=$1 and deleted_at is null`, idStr)

	if err != nil {
		return err
	}

	n, err := rows.RowsAffected()

	if err != nil {
		return err
	}

	if n > 0 {
		return nil
	}

	return errors.New("Cannot delete Author becouse Author not found")
}

// hard delete uchun kod
// func (stg Postgres) removeAuthorDelete(slice []models.Author, s int) []models.Author {
// 	return append(slice[:s], slice[s+1:]...)
// }
