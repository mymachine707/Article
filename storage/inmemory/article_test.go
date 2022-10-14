package inmemory_test

import (
	"errors"
	"mymachine707/models"
	"mymachine707/storage/inmemory"
	"testing"
)

func TestArticle(t *testing.T) {
	var err error
	IM := inmemory.InMemory{
		Db: &inmemory.DB{},
	}

	expextedErrorAuthor := errors.New("author not found")
	authorId := "b3546729-0695-4c63-ba3d-c3caa7310cde"
	authorData := models.CreateAuthorModul{
		Firstname: "John",
		Lastname:  "Doe",
	}
	NotFoundAuthorId := "249d62ba-b898-435b-b35e-ad7e505fc604"

	contents := models.Content{
		Title: "Lorem",
		Body:  "Impsum",
	}

	err = IM.AddAuthor(authorId, authorData)

	if err != nil {
		t.Fatalf("unexpextedError: %v", err)
	}

	var TestAddArticle = []struct {
		name       string
		id         string
		data       models.CreateArticleModul
		wantError  error
		wantResult models.CreateArticleModul
	}{
		{
			name: "success",
			id:   "836e2951-8190-40b3-8d02-3a6a6b34f4a5",
			data: models.CreateArticleModul{
				Content:  contents,
				AuthorID: authorId,
			},
			wantError: nil,
			wantResult: models.CreateArticleModul{
				Content: contents,
			},
		},
		{
			name: "fail",
			id:   "836e2951-8190-40b3-8d02-3a6a6b34f4a5",
			data: models.CreateArticleModul{
				Content:  contents,
				AuthorID: NotFoundAuthorId,
			},
			wantError:  expextedErrorAuthor,
			wantResult: models.CreateArticleModul{},
		},
	}

	for _, v := range TestAddArticle {
		t.Run(v.name, func(t *testing.T) {

			err := IM.AddArticle(v.id, v.data)

			if v.wantError == nil {
				if err != nil {
					t.Errorf("unexpexted Error: %v", err)
				}
				article, err := IM.GetArticleByID(v.id)

				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}

				if v.wantResult.Content != article.Content {
					t.Errorf("We want result: %v but got %v", v.wantResult.Content, article.Content)
				}
			} else {
				if v.wantError.Error() != err.Error() {
					t.Errorf("We want error: %v but got error: %v", v.wantError, err)
				}

			}
		})
	}

	articleId := "836e2951-8190-40b3-8d02-3a6a6b34f4a5"
	articleData := models.CreateArticleModul{
		Content: models.Content{
			Title: "Lorem",
			Body:  "Impsum",
		},
		AuthorID: "b3546729-0695-4c63-ba3d-c3caa7310cde",
	}

	err = IM.AddArticle(articleId, articleData)
	if err != nil && err.Error() != expextedErrorAuthor.Error() {
		t.Errorf("IM.AddArticle() expexted: %v, but got error: %v", expextedErrorAuthor, err)
	}

	if err != nil && err.Error() != expextedErrorAuthor.Error() {
		t.Errorf("IM.AddArticle() expexted: %v, but got error: %v", expextedErrorAuthor, err)
	}

	//

	t.Log("Test has been finished!")
	// go test -coverprofile=coverage.out ./... ;    go tool cover -html=coverage.out
}

//<---------------------------- my code ---------------------------------------------->
/*
err = IM.AddAuthor("b3546729-0695-4c63-ba3d-c3caa7310cde", models.CreateAuthorModul{
	Firstname: "John",
	Lastname:  "Doe",
})

if err != nil {
	t.Fatalf("unexpected err: %v", err)
}

TestAddArticle := []struct {
	name        string
	id          string
	entity      models.CreateArticleModul
	wantedError error
}{
	{
		name: "TestAddArticle",
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
		name: "TestAddArticle",
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
		name: "TestAddArticle",
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

for _, v := range TestAddArticle {
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
*/
//<---------------------------- my code ---------------------------------------------->
// 	//func (IM InMemory) GetArticleByID(id string) (models.PackedArticleModel, error)
// 	//author, err := IM.GetAuthorByID(entity.AuthorID)

// 	TestGetArticleByID := []struct {
// 		name        string
// 		id          string
// 		result      models.PackedArticleModel
// 		wantedError error
// 	}{

// 		{
// 			name: "TestGetArticleById",
// 			id:   "19ecfdb8-5f01-4a36-805a-ff8bc34c59ab",
// 			result: models.PackedArticleModel{
// 				ID: "19ecfdb8-5f01-4a36-805a-ff8bc34c59ab",
// 				Content: models.Content{
// 					Title: "Lorem",
// 					Body:  "Impsum",
// 				},
// 				Author: models.Author{
// 					ID:        "b3546729-0695-4c63-ba3d-c3caa7310cde",
// 					Firstname: "John",
// 					Lastname:  "Doe",
// 				},
// 			},
// 			wantedError: nil,
// 		},
// 	}

// 	for _, v := range TestGetArticleByID {
// 		t.Run(v.name, func(t *testing.T) {

// 		article, err = IM.GetArticleByID(v.id)

// 		if article.

// 		if err!=nil {

// 		}

// 			if err != nil && v.wantedError != nil {
// 				if err.Error() != v.wantedError.Error() {
// 					t.Fatalf("we wanted error-->: %v, but we got-->: %v", v.wantedError, err)
// 				}
// 			}

// 			if (err != nil && v.wantedError == nil) || (err == nil && v.wantedError != nil) {
// 				t.Fatalf("-----------------------------> Somthing went wrong! <------------------------------------\n error: %v", err)
// 			}
// 		})

// 	}
// //<----------------------------------------- end test get article by id -------------------------------------->
