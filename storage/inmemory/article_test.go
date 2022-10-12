package inmemory_test

import (
	"errors"
	"mymachine707/models"
	"mymachine707/storage/inmemory"
	"testing"
)

func TestAddArticle(t *testing.T) {

	var err error
	IM := inmemory.InMemory{
		Db: &inmemory.DB{},
	}

	if err != nil {
		t.Fatalf("unexpextedError: %v", err)
	}

	expextedError := errors.New("author not found")

	if err != nil && err.Error() != expextedError.Error() {
		t.Errorf("IM.AddArticle() expexted: %v, but got error: %v", expextedError, err)
	}

	err = IM.AddAuthor("b3546729-0695-4c63-ba3d-c3caa7310cde", models.CreateAuthorModul{
		Firstname: "John",
		Lastname:  "Doe",
	})

	err = IM.AddArticle("836e2951-8190-40b3-8d02-3a6a6b34f4a5", models.CreateArticleModul{
		Content: models.Content{
			Title: "Lorem",
			Body:  "Impsum",
		},
		AuthorID: "b3546729-0695-4c63-ba3d-c3caa7310cde",
	})

	if err != nil && err.Error() != expextedError.Error() {
		t.Errorf("IM.AddArticle() expexted: %v, but got error: %v", expextedError, err)
	}

	t.Log("Test has been finished!")
}
