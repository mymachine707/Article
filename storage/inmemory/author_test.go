package inmemory_test

import (
	"errors"
	"mymachine707/models"
	"mymachine707/storage/inmemory"
	"testing"
)

func TestAddAuthor(t *testing.T) {
	var err error
	IM := inmemory.InMemory{
		Db: &inmemory.DB{},
	}

	//AddAuthor(id string, entity models.CreateAuthorModul) error

	var TestAddAuthor = []struct {
		name        string
		id          string
		data        models.CreateAuthorModul
		wantedError error
	}{
		{
			name: "success",
			id:   "fa54ef72-76e6-4368-b772-117e7bb9ca01",
			data: models.CreateAuthorModul{
				Firstname: "John",
				Lastname:  "doe",
			},
			wantedError: nil,
		},
		{
			name: "fail: id must exist",
			id:   "",
			data: models.CreateAuthorModul{
				Firstname: "John",
				Lastname:  "doe",
			},
			wantedError: errors.New("id must exist"),
		},
		{
			name: "fail: Firstname must exist",
			id:   "fa54ef72-76e6-4368-b772-117e7bb9ca01",
			data: models.CreateAuthorModul{
				Firstname: "",
				Lastname:  "doe",
			},
			wantedError: errors.New("Firstname must exist"),
		},
		{
			name: "fail: Lastname must exist",
			id:   "fa54ef72-76e6-4368-b772-117e7bb9ca01",
			data: models.CreateAuthorModul{
				Firstname: "John",
				Lastname:  "",
			},
			wantedError: errors.New("Lastname must exist"),
		},
	}

	for _, v := range TestAddAuthor {
		t.Run(v.name, func(t *testing.T) {
			err = IM.AddAuthor(v.id, v.data)
			if v.wantedError == nil {
				if err != nil {
					t.Errorf("Unexpexted error: %v", err)
				}
			} else {
				if err != nil && v.wantedError.Error() != err.Error() {
					t.Errorf("unexpexted error: %v", err)
				}
			}
		})
	}
	//
	t.Log("<----------------- Auhtor test finished ----------------->")
	// go test -coverprofile=coverage.out ./... ;    go tool cover -html=coverage.out
}
func TestGetAuthorByID(t *testing.T) {
	var err error
	IM := inmemory.InMemory{
		Db: &inmemory.DB{},
	}

	err = IM.AddAuthor("c3bc996a-e2a1-4143-85da-7c4fc378dca5", models.CreateAuthorModul{
		Firstname: "John",
		Lastname:  "Doe",
	})

	if err != nil {
		t.Errorf("Unexpexted error: %v", err)
	}

	// deleted teskshirish uchun
	err = IM.AddAuthor("20b50128-7575-4b28-9eb7-bb731f7caf16", models.CreateAuthorModul{
		Firstname: "Mikel",
		Lastname:  "Mackdon",
	})

	if err != nil {
		t.Errorf("Unexpexted error: %v", err)
	}

	err = IM.DeleteAuthor("20b50128-7575-4b28-9eb7-bb731f7caf16")

	if err != nil {
		t.Errorf("Unexpexted error: %v", err)
	}

	//AddAuthor(id string, entity models.CreateAuthorModul) error

	var TestGetAuthorByID = []struct {
		name        string
		id          string
		wantResult  models.Author
		wantedError error
	}{
		{
			name: "success",
			id:   "c3bc996a-e2a1-4143-85da-7c4fc378dca5",
			wantResult: models.Author{
				ID:        "c3bc996a-e2a1-4143-85da-7c4fc378dca5",
				Firstname: "John",
				Lastname:  "Doe",
			},
			wantedError: nil,
		},
		{
			name:        "Fail: author alredy deleted",
			id:          "20b50128-7575-4b28-9eb7-bb731f7caf16",
			wantResult:  models.Author{},
			wantedError: errors.New("author already deleted"),
		},
	}

	for _, v := range TestGetAuthorByID {
		t.Run(v.name, func(t *testing.T) {

			author, err := IM.GetAuthorByID(v.id)

			if v.wantedError == nil {
				if err != nil {
					t.Errorf("Unexpexted error: %v", err)
				}
			} else {
				for _, s := range IM.Db.InMemoryAuthorData {
					if s.ID == v.id && s.DeletedAt != nil {
						if v.wantedError.Error() != err.Error() {
							t.Errorf("We wanted error: %v, bot got error:%v", v.wantedError, err)
						}
					}
				}
				if err != nil && (v.wantResult.Firstname != author.Firstname || v.wantResult.Lastname != author.Lastname) {
					t.Errorf("unexpexted error: %v", err)
				}

			}
		})
	}
	//
	t.Log("<----------------- Auhtor test finished ----------------->")
	// go test -coverprofile=coverage.out ./... ;    go tool cover -html=coverage.out
}

func TestAuthorUpdate(t *testing.T) {
	var err error

	IM := inmemory.InMemory{
		Db: &inmemory.DB{},
	}

	err = IM.AddAuthor("c3bc996a-e2a1-4143-85da-7c4fc378dca5", models.CreateAuthorModul{
		Firstname: "John",
		Lastname:  "Doe",
	})

	if err != nil {
		t.Errorf("Unexpexted error: %v", err)
	}

	//func (IM InMemory) UpdateAuthor(author models.UpdateAuthorModul) error {
	var TestAuthorUpdate = []struct {
		name      string
		author    models.UpdateAuthorModul
		wantError error
	}{
		{
			name: "success",
			author: models.UpdateAuthorModul{
				ID:        "c3bc996a-e2a1-4143-85da-7c4fc378dca5",
				Firstname: "John",
				Lastname:  "Doe",
			},
			wantError: nil,
		},
		{
			name: "fail author not found",
			author: models.UpdateAuthorModul{
				ID:        "5e86359c-80ec-4bd0-adbc-06aad8999548",
				Firstname: "Mike",
				Lastname:  "Tayson",
			},
			wantError: errors.New("author not found"),
		},
	}

	for _, v := range TestAuthorUpdate {
		t.Run(v.name, func(t *testing.T) {

			err = IM.UpdateAuthor(v.author)

			if v.wantError == nil {
				if err != nil {
					t.Errorf("Unexpexted error: %v", err)
				}
			} else {
				if err != nil && v.wantError.Error() != err.Error() {
					t.Errorf("We want err: %v, but got err: %v", v.wantError, err)
				}
			}
		})
	}
	t.Log("<----------------- Auhtor test finished ----------------->")
	// go test -coverprofile=coverage.out ./... ;    go tool cover -html=coverage.out
}

func TestGetAuthorList(t *testing.T) {
	var err error

	IM := inmemory.InMemory{
		Db: &inmemory.DB{},
	}

	err = IM.AddAuthor("c3bc996a-e2a1-4143-85da-7c4fc378dca5", models.CreateAuthorModul{
		Firstname: "John",
		Lastname:  "Doe",
	})

	if err != nil {
		t.Errorf("Unexpexted error: %v", err)
	}

	err = IM.AddAuthor("121ec437-9a7b-482a-adf4-9f73f21bbbe4", models.CreateAuthorModul{
		Firstname: "Mike",
		Lastname:  "Tayson",
	})

	if err != nil {
		t.Errorf("Unexpexted error: %v", err)
	}

	err = IM.AddAuthor("1f594cbf-1bf7-45f4-b7b8-8105e0112e08", models.CreateAuthorModul{
		Firstname: "Mike",
		Lastname:  "Tayson",
	})

	if err != nil {
		t.Errorf("Unexpexted error: %v", err)
	}

	// func (IM InMemory) GetAuthorList(offset, limit int, serach string) (resp []models.Author, err error)
	var TestGetAuthorList = []struct {
		name       string
		offset     int
		limit      int
		search     string
		wantResult []models.Author
		wantError  error
	}{
		{
			name:   "success",
			offset: 1,
			limit:  2,
			search: "Lorem",
			wantResult: []models.Author{
				{
					ID:        "121ec437-9a7b-482a-adf4-9f73f21bbbe4",
					Firstname: "Mike",
					Lastname:  "Tayson",
				},
				{
					ID:        "c3bc996a-e2a1-4143-85da-7c4fc378dca5",
					Firstname: "John",
					Lastname:  "Doe",
				},
				{
					ID:        "1f594cbf-1bf7-45f4-b7b8-8105e0112e08",
					Firstname: "John",
					Lastname:  "Doe",
				},
			},
			wantError: nil,
		},
	}

	for _, v := range TestGetAuthorList {
		t.Run(v.name, func(t *testing.T) {
			authorList, err := IM.GetArticleList(v.offset, v.limit, v.search)
			if v.wantError == nil {
				if err != nil {
					t.Errorf("Unexpexted error^ %v", err)
				}
			} else {
				for _, s := range authorList {
					if s.DeletedAt != nil {
						t.Errorf("author already deleted")
					}
				}

				if err == nil {
					t.Errorf("unexpexted error")
				}

			}
			// } else {
			// 	if
			// }

		})
	}
	t.Log("<----------------- Auhtor test finished ----------------->")
	// go test -coverprofile=coverage.out ./... ;    go tool cover -html=coverage.out
}
