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
