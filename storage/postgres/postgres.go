package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Postgres ...
type Postgres struct {
	db *sqlx.DB
}

var schema = `
CREATE TABLE IF NOT EXISTS article (
   id CHAR(36) PRIMARY KEY,
   title VARCHAR(255) UNIQUE NOT NULL,
   body TEXT NOT NULL,
   author_id CHAR(36),
   created_at TIMESTAMP DEFAULT now(),
   updated_at TIMESTAMP,
   deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS author (
	id CHAR(36) PRIMARY KEY,
	firstname VARCHAR(255) NOT NULL,
	lastname VARCHAR(255) NOT NULL,
	created_at TIMESTAMP DEFAULT now(),
	updated_at TIMESTAMP,
	deleted_at TIMESTAMP
 );


ALTER TABLE article DROP CONSTRAINT IF EXISTS fk_article_author;
ALTER TABLE article ADD CONSTRAINT fk_article_author FOREIGN KEY (author_id) REFERENCES author (id);
 `

// InitDB ...
func InitDB(psqlConfig string) (*Postgres, error) {
	var err error

	tempDB, err := sqlx.Connect("postgres", psqlConfig)
	if err != nil {
		return nil, err
	}

	tempDB.MustExec(schema)

	tx := tempDB.MustBegin()

	tx.MustExec("INSERT INTO author (id, firstname, lastname) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING", "b9401ecc-e7b7-4e83-b387-eb85072adcd9", "John", "Doe")
	tx.MustExec("INSERT INTO author (id, firstname, lastname) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING", "1f27a12d-93c7-4272-9eec-43e28a00482d", "Jason", "Moiron")

	tx.MustExec("INSERT INTO article (id, title, body, author_id) VALUES ($1, $2, $3, $4) ON CONFLICT DO NOTHING", "3d5ee64f-1810-404f-a804-58f12dd18279", "Lorem 1", "Body 1", "1f27a12d-93c7-4272-9eec-43e28a00482d")
	tx.MustExec("INSERT INTO article (id, title, body, author_id) VALUES ($1, $2, $3, $4) ON CONFLICT DO NOTHING", "cc9be5d9-aa7b-48a7-9bd5-737fd48f37f0", "Lorem 2", "Body 2", "b9401ecc-e7b7-4e83-b387-eb85072adcd9")

	tx.NamedExec("INSERT INTO article (id, title, body, author_id) VALUES (:id, :t, :b, :aid) ON CONFLICT DO NOTHING", map[string]interface{}{
		"id":  "80f20849-b8a6-4c4e-b589-c9511b4145c4",
		"t":   "Lorem 3",
		"b":   "Body 3",
		"aid": "b9401ecc-e7b7-4e83-b387-eb85072adcd9",
	})
	err = tx.Commit()

	if err != nil {
		return nil, err
	}

	return &Postgres{
		db: tempDB,
	}, nil
}
