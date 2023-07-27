package golangtest

import (
	"encoding/json"
	"net/http"
)

type UserDb interface {
	FindById(ID string) (string, error)
}

func GetUser(db UserDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type User struct {
			Name string `json:"name"`
		}

		name, err := db.FindById("1")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		user1 := User{
			Name: name,
		}

		b, _ := json.Marshal(user1)
		w.Write(b)
		return
	}
}
