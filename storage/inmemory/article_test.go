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
		{
			name: "fail: id must exist",
			id:   "",
			data: models.CreateArticleModul{
				Content:  contents,
				AuthorID: NotFoundAuthorId,
			},
			wantError:  errors.New("id must exist"),
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

	t.Log("<----------------- Article test finished ----------------->")
	// go test -coverprofile=coverage.out ./... ;    go tool cover -html=coverage.out
}
func TestGetArticleById(t *testing.T) {
	var err error
	IM := inmemory.InMemory{
		Db: &inmemory.DB{},
	}

	authorId := "b3546729-0695-4c63-ba3d-c3caa7310cde"
	authorData := models.CreateAuthorModul{
		Firstname: "John",
		Lastname:  "Doe",
	}
	err = IM.AddAuthor(authorId, authorData)

	if err != nil {
		t.Fatalf("unexpextedError: %v", err)
	}
	contents := models.Content{
		Title: "Lorem",
		Body:  "Impsum",
	}

	err = IM.AddArticle("249d62ba-b898-435b-b35e-ad7e505fc604", models.CreateArticleModul{
		Content:  contents,
		AuthorID: authorId,
	})

	if err != nil {
		t.Fatalf("unexpextedError: %v", err)
	}
	// deleted tekshirish uchun
	err = IM.AddArticle("de36b75b-d496-40fa-9d7c-28183930b3a6", models.CreateArticleModul{
		Content: models.Content{
			Title: "deleted uchun",
			Body:  "deleted uchun",
		},
		AuthorID: authorId,
	})

	if err != nil {
		t.Fatalf("unexpextedError: %v", err)
	}

	err = IM.DeleteArticle("de36b75b-d496-40fa-9d7c-28183930b3a6")

	if err != nil {
		t.Fatalf("unexpextedError: %v", err)
	}

	var TestGetAddArticleByID = []struct {
		name       string
		id         string
		wantError  error
		wantResult models.PackedArticleModel
	}{
		{
			name:      "success",
			id:        "249d62ba-b898-435b-b35e-ad7e505fc604",
			wantError: nil,
			wantResult: models.PackedArticleModel{
				ID:      "249d62ba-b898-435b-b35e-ad7e505fc604",
				Content: contents,
				Author: models.Author{
					ID:        authorId,
					Firstname: "John",
					Lastname:  "Doe",
				},
			},
		},
		{
			name:      "fail: id must exist",
			id:        "",
			wantError: errors.New("id must exist"),
			wantResult: models.PackedArticleModel{
				ID:      "249d62ba-b898-435b-b35e-ad7e505fc604",
				Content: contents,
				Author: models.Author{
					ID:        authorId,
					Firstname: "John",
					Lastname:  "Doe",
				},
			},
		},
		{
			name:      "fail: article not found",
			id:        "42056f47-9f1b-4a40-ad9c-0930f7faaf1f",
			wantError: errors.New("article not found"),
			wantResult: models.PackedArticleModel{
				ID:      "249d62ba-b898-435b-b35e-ad7e505fc604",
				Content: contents,
				Author: models.Author{
					ID:        authorId,
					Firstname: "John",
					Lastname:  "Doe",
				},
			},
		},
		{
			name:       "fail: article already deleted",
			id:         "de36b75b-d496-40fa-9d7c-28183930b3a6",
			wantError:  errors.New("article already deleted"),
			wantResult: models.PackedArticleModel{},
		},
	}

	for _, v := range TestGetAddArticleByID {
		t.Run(v.name, func(t *testing.T) {

			article, err := IM.GetArticleByID(v.id)

			if v.wantError == nil {
				if err != nil {
					t.Errorf("unexpexted Error: %v", err)
				}

				if v.wantResult.Content != article.Content {
					t.Errorf("We want result content: %v but got %v", v.wantResult.Content, article.Content)
				}
			} else {

				for _, s := range IM.Db.InMemoryArticleData {
					if s.ID == v.id && s.DeletedAt != nil {
						if err.Error() != v.wantError.Error() {
							t.Errorf("Unexpexted error: %v", err)
						}
					}
				}
				if v.id != v.wantResult.ID {
					if v.wantError.Error() != err.Error() {
						t.Errorf("Method: %v, article not found", v.name)
					}
				}
				if v.wantError.Error() != err.Error() {
					t.Errorf("We want error: %v but got error: %v", v.wantError, err)
				}
				if err == nil {
					t.Errorf("unexpexted error")
				}

			}
		})
	}

	//

	t.Log("<----------------- Article test finished ----------------->")
	// go test -coverprofile=coverage.out ./... ;    go tool cover -html=coverage.out
}
func TestGetArticleList(t *testing.T) {
	var err error
	IM := inmemory.InMemory{
		Db: &inmemory.DB{},
	}

	authorId := "b3546729-0695-4c63-ba3d-c3caa7310cde"
	authorData := models.CreateAuthorModul{
		Firstname: "John",
		Lastname:  "Doe",
	}
	err = IM.AddAuthor(authorId, authorData)

	if err != nil {
		t.Fatalf("unexpextedError: %v", err)
	}

	contents := models.Content{
		Title: "Lorem",
		Body:  "Impsum",
	}
	// article list
	//0
	err = IM.AddArticle("249d62ba-b898-435b-b35e-ad7e505fc604", models.CreateArticleModul{
		Content:  contents,
		AuthorID: authorId,
	})

	if err != nil {
		t.Fatalf("unexpextedError: %v", err)
	}
	// 1
	err = IM.AddArticle("249d62b1-b898-435b-b35e-ad7e505fc604", models.CreateArticleModul{
		Content:  contents,
		AuthorID: authorId,
	})

	if err != nil {
		t.Fatalf("unexpextedError: %v", err)
	}

	var TestGetAddArticleList = []struct {
		name       string
		offset     int
		limit      int
		search     string
		wantError  error
		wantResult []models.Article
	}{
		{
			name:      "success",
			offset:    1,
			limit:     1,
			search:    "Lorem",
			wantError: nil,
			wantResult: []models.Article{
				{
					ID:        "249d62ba-b898-435b-b35e-ad7e505fc604",
					Content:   contents,
					AuthorID:  authorId,
					DeletedAt: nil,
				},
				{
					ID:        "249d62b1-b898-435b-b35e-ad7e505fc604",
					Content:   contents,
					AuthorID:  authorId,
					DeletedAt: nil,
				},
			},
		},
	}

	for _, v := range TestGetAddArticleList {
		t.Run(v.name, func(t *testing.T) {

			article, err := IM.GetArticleList(v.offset, v.limit, v.search)

			if v.wantError == nil {
				if err != nil {
					t.Errorf("unexpexted Error: %v", err)
				}
			} else {

				for _, s := range article {
					if s.DeletedAt != nil {
						t.Errorf("article already deleted")
					}
				}

				if err == nil {
					t.Errorf("unexpexted error")
				}

			}
		})
	}
	t.Log("<----------------- Article test finished ----------------->")
	// go test -coverprofile=coverage.out ./... ;    go tool cover -html=coverage.out
}

func TestUpdateArticle(t *testing.T) {
	var err error
	IM := inmemory.InMemory{
		Db: &inmemory.DB{},
	}

	authorId := "b3546729-0695-4c63-ba3d-c3caa7310cde"
	authorData := models.CreateAuthorModul{
		Firstname: "John",
		Lastname:  "Doe",
	}
	err = IM.AddAuthor(authorId, authorData)

	if err != nil {
		t.Fatalf("unexpextedError: %v", err)
	}

	contents := models.Content{
		Title: "Lorem",
		Body:  "Impsum",
	}
	// article list
	//0
	err = IM.AddArticle("249d62ba-b898-435b-b35e-ad7e505fc604", models.CreateArticleModul{
		Content:  contents,
		AuthorID: authorId,
	})

	if err != nil {
		t.Fatalf("unexpextedError: %v", err)
	}

	var TestUpdateArticle = []struct {
		name      string
		article   models.UpdateArticleModul
		wantError error
	}{
		{
			name: "success",
			article: models.UpdateArticleModul{
				ID:      "249d62ba-b898-435b-b35e-ad7e505fc604",
				Content: contents,
			},
			wantError: nil,
		},
		{
			name: "fail: article not found",
			article: models.UpdateArticleModul{
				ID:      "25459a1c-0511-4f11-b5c0-9d2f8ccb7e8f",
				Content: contents,
			},
			wantError: errors.New("article not found"),
		},
	}

	for _, v := range TestUpdateArticle {
		t.Run(v.name, func(t *testing.T) {

			err := IM.UpdateArticle(v.article)

			if v.wantError == nil {
				if err != nil {
					t.Errorf("unexpexted Error: %v", err)
				}
			} else {
				if err != nil && v.wantError.Error() != err.Error() {
					t.Errorf("We want error:%v, but got error: %v", v.wantError, err)
				}
				if err == nil {
					t.Errorf("unexpexted error")
				}
			}
		})
	}
	t.Log("<----------------- Article test finished ----------------->")
	// go test -coverprofile=coverage.out ./... ;    go tool cover -html=coverage.out
}

func TestDeleteArticle(t *testing.T) {
	var err error
	IM := inmemory.InMemory{
		Db: &inmemory.DB{},
	}

	authorId := "b3546729-0695-4c63-ba3d-c3caa7310cde"
	authorData := models.CreateAuthorModul{
		Firstname: "John",
		Lastname:  "Doe",
	}
	err = IM.AddAuthor(authorId, authorData)

	if err != nil {
		t.Fatalf("unexpextedError: %v", err)
	}

	err = IM.AddArticle("a6ed6b07-9239-4cd5-9c24-cbb6b596d438", models.CreateArticleModul{
		Content: models.Content{
			Title: "deleted uchun",
			Body:  "deleted uchun",
		},
		AuthorID: authorId,
	})

	if err != nil {
		t.Fatalf("unexpextedError: %v", err)
	}

	// deleted tekshirish uchun
	err = IM.AddArticle("de36b75b-d496-40fa-9d7c-28183930b3a6", models.CreateArticleModul{
		Content: models.Content{
			Title: "deleted uchun",
			Body:  "deleted uchun",
		},
		AuthorID: authorId,
	})

	if err != nil {
		t.Fatalf("unexpextedError: %v", err)
	}

	err = IM.DeleteArticle("de36b75b-d496-40fa-9d7c-28183930b3a6")

	if err != nil {
		t.Fatalf("unexpextedError: %v", err)
	}

	var TestDeleteArticle = []struct {
		name      string
		id        string
		wantError error
	}{
		{
			name:      "success",
			id:        "a6ed6b07-9239-4cd5-9c24-cbb6b596d438",
			wantError: nil,
		},
		{
			name:      "fail: article already deleted",
			id:        "a6ed6b07-9239-4cd5-9c24-cbb6b596d438",
			wantError: errors.New("article already deleted"),
		},
		{
			name:      "fail: article not found",
			id:        "938edc67-df1d-465c-bb25-0536a66239be",
			wantError: errors.New("Cannot delete article becouse Article not found"),
		},
	}

	for _, v := range TestDeleteArticle {
		t.Run(v.name, func(t *testing.T) {

			err := IM.DeleteArticle(v.id)

			if v.wantError == nil {
				if err != nil {
					t.Errorf("unexpexted Error: %v", err)
				}
			} else {
				for _, s := range IM.Db.InMemoryArticleData {
					if s.ID == v.id && s.DeletedAt != nil {
						if err.Error() != v.wantError.Error() {
							t.Errorf("Unexpexted error: %v", err)
						}
					}
				}
				if err != nil && v.wantError.Error() != err.Error() {
					t.Errorf("We want error: %v, but got: %v", v.wantError, err)
				}
				if err == nil {
					t.Errorf("unexpexted error")
				}

			}
		})
	}

	//

	t.Log("<----------------- Article test finished ----------------->")
	// go test -coverprofile=coverage.out ./... ;    go tool cover -html=coverage.out
}

//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//

//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//

//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//

//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//

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
