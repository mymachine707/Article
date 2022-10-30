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
	$1,
	$2
	$3,
	#4
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

// GetArticleByID ...
func (stg Postgres) GetArticleByID(id string) (models.PackedArticleModel, error) {
	var a models.PackedArticleModel
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
    au.created_at,
    au.updated_at,
    au.deleted_at
 FROM article AS ar JOIN author AS au ON ar.author_id = au.id WHERE ar.id = $1;`, id).Scan(
		&a.ID,
		&a.Title,
		&a.Body,
		&a.CreatedAt,
		&a.UpdatedAt,
		&a.DeletedAt,
		&a.Author.ID,
		&a.Author.Firstname,
		&a.Author.Lastname,
		&a.Author.CreatedAt,
		&a.Author.UpdatedAt,
		&a.Author.DeletedAt,
	)

	if err != nil {
		return a, err
	}
	return a, nil
}

// GetArticleList ...
func (stg Postgres) GetArticleList(offset, limit int, search string) (resp []models.Article, err error) {

	rows, err := stg.db.Queryx("SELECT * FROM article")
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

	// for i, v := range IM.Db.InMemoryArticleData {
	// 	if v.ID == article.ID && v.DeletedAt == nil {

	// 		v.Content = article.Content
	// 		t := time.Now()
	// 		v.UpdatedAt = &t

	// 		IM.Db.InMemoryArticleData[i] = v

	// 		return nil
	// 	}
	// }
	return errors.New("article not found")
}

// DeleteArticle ...
func (stg Postgres) DeleteArticle(idStr string) error {

	// for i, v := range IM.Db.InMemoryArticleData {
	// 	if v.ID == idStr {
	// 		if v.DeletedAt != nil {
	// 			return errors.New("article already deleted")
	// 		}
	// 		// bu kod article hard delete qilish uchun :
	// 		// IM.Db.InMemoryArticleData = remove(IM.Db.InMemoryArticleData, i)

	// 		// bu kod soft delete uchun:
	// 		t := time.Now()
	// 		v.DeletedAt = &t
	// 		IM.Db.InMemoryArticleData[i] = v
	// 		return nil
	// 	}
	// }
	return errors.New("Cannot delete article becouse Article not found")
}

// hard delete uchun kod
// func (IM InMemory) remove(slice []models.Article, s int) []models.Article {
// 	return append(slice[:s], slice[s+1:]...)
// }
