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
	/*
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
	*/

	err = IM.AddAuthor("b3546729-0695-4c63-ba3d-c3caa7310cde", models.CreateAuthorModul{
		Firstname: "John",
		Lastname:  "Doe",
	})

	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}

	articletest := []struct {
		name        string
		id          string
		entity      models.CreateArticleModul
		wantedError error
	}{
		{
			name: "successful",
			id:   "19ecfdb8-5f01-4a36-805a-ff8bc34c59ab",
			entity: models.CreateArticleModul{
				Content: models.Content{
					Title: "Lorem",
					Body:  "impsum",
				},
				//AuthorID: "b3546729-0695-4c63-ba3d-c3caa7310cde",
			},
			wantedError: errors.New("author not found"),
		},
		{
			name: "successful",
			id:   "",
			entity: models.CreateArticleModul{
				Content: models.Content{
					Title: "Lorem",
					Body:  "impsum",
				},
				AuthorID: "b3546729-0695-4c63-ba3d-c3caa7310cde",
			},
			wantedError: errors.New("id must exist"),
		},

		{
			name: "successful",
			id:   "19ecfdb8-5f01-4a36-805a-ff8bc34c59ab",
			entity: models.CreateArticleModul{
				Content: models.Content{
					Title: "Lorem",
					Body:  "impsum",
				},
				AuthorID: "b3546729-0695-4c63-ba3d-c3caa7310cde",
			},
			wantedError: nil,
		},
	}

	for _, v := range articletest {
		t.Run(v.name, func(t *testing.T) {

			err = IM.AddArticle(v.id, v.entity)

			if err != nil && v.wantedError != nil {
				if err.Error() != v.wantedError.Error() {
					t.Fatalf("we wanted error-->: %v, but we got-->: %v", v.wantedError, err)
				}
			}

			if (err != nil && v.wantedError == nil) || (err == nil && v.wantedError != nil) {
				t.Fatalf("-----------------------------> Somthing went wrong! <------------------------------------\n error: %v", err)
			}
		})

	}

	t.Log("Test has been finished!")
}
