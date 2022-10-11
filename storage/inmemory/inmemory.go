package inmemory

import "mymachine707/models"

// InMemory ...
type InMemory struct {
	Db *DB
}

// DB mock ... ???
type DB struct {
	// InMemoryAuthorData - data base author
	InMemoryAuthorData []models.Author

	// InMemoryArticleData - data base article
	InMemoryArticleData []models.Article
}
