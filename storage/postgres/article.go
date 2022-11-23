package postgres

import (
	"errors"
	"log"
	"mymachine707/models"
)

// AddArticle ...
func (stg Postgres) AddArticle(id string, entity models.CreateArticleModul) error {
	if id == "" {
		return errors.New("id must exist")
	}

	_, err := stg.GetAuthorByID(entity.AuthorID)

	if err != nil {
		return err
	}

	_, err = stg.db.Exec(`INSERT INTO article (
		id,
		title,
		body,
		author_id		
		) VALUES (
		$1,
		$2,
		$3,
		$4
)`,
		id,
		entity.Title,
		entity.Body,
		entity.AuthorID,
	)

	if err != nil {
		return err
	}

	return nil
}

// GetArticleByID ...  //  ????
func (stg Postgres) GetArticleByID(id string) (models.PackedArticleModel, error) {
	var a models.PackedArticleModel

	var tempMiddlename *string

	if id == "" {
		return a, errors.New("id must exist")
	}

	err := stg.db.QueryRow(`SELECT 
    ar.id,
    ar.title,
    ar.body,
    ar.created_at,
    ar.updated_at,
    ar.deleted_at,
    au.id,
    au.firstname,
    au.lastname,
	au.middlename,
	au.fullname,
    au.created_at,
    au.updated_at,
    au.deleted_at FROM article AS ar JOIN author AS au ON ar.author_id = au.id WHERE ar.id = $1`, id).Scan(
		&a.ID,
		&a.Title,
		&a.Body,
		&a.CreatedAt,
		&a.UpdatedAt,
		&a.DeletedAt,
		&a.Author.ID,
		&a.Author.Firstname,
		&a.Author.Lastname,
		&tempMiddlename,
		&a.Author.Fullname,
		&a.Author.CreatedAt,
		&a.Author.UpdatedAt,
		&a.Author.DeletedAt,
	)

	if tempMiddlename != nil {
		a.Author.Middlename = *tempMiddlename
	}

	if err != nil {
		return a, err
	}

	return a, nil
}

// GetArticleList ...
func (stg Postgres) GetArticleList(offset, limit int, search string) (resp []models.Article, err error) {

	rows, err := stg.db.Queryx(`
	
	SELECT * FROM article WHERE 
	
	(title || ' ' || body ILIKE '%' || $1 || '%') AND deleted_at is null
	LIMIT $2
	OFFSET $3
	
	`, search, limit, offset)
	
	
	if err != nil {
		return resp, err
	}

	for rows.Next() {
		var a models.Article
		err := rows.Scan(
			&a.ID,
			&a.Title,
			&a.Body,
			&a.AuthorID,
			&a.CreatedAt,
			&a.UpdatedAt,
			&a.DeletedAt,
		)

		//err := rows.StructScan(&a)
		if err != nil {
			log.Panic(err)
		}

		//fmt.Printf("%d a---> %#v\n", i, a)
		resp = append(resp, a)
	}

	return resp, err
}

// UpdateArticle ...
func (stg Postgres) UpdateArticle(article models.UpdateArticleModul) error {

	res, err := stg.db.NamedExec("UPDATE article SET title=:t, body=:b, updated_at=now() WHERE id=:id AND deleted_at is null", map[string]interface{}{
		"id": article.ID,
		"t":  article.Title,
		"b":  article.Body,
	})

	if err != nil {
		return err
	}

	n, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if n > 0 {
		return nil
	}

	return errors.New("article not found")
}

// DeleteArticle ...
func (stg Postgres) DeleteArticle(idStr string) error {

	res, err := stg.db.Exec("UPDATE article Set deleted_at=now() WHERE id=$1 AND deleted_at is null", idStr)

	if err != nil {
		return err
	}

	n, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if n > 0 {
		return nil
	}

	return errors.New("article not found")
}
