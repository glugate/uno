package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/glugate/uno/apps/crm/repo"
	"github.com/glugate/uno/pkg/uno/server"
)

// UsersFind returns one user
func UsersFind(w http.ResponseWriter, r *http.Request) {
	id := server.RequestParam(r, 0)
	repo, err := repo.NewUserRepo()
	if err != nil {
		panic(err)
	}
	items, err := repo.Find(id)
	if err != nil {
		panic(err)
	}

	// Get the user items

	out, err := json.Marshal(items)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write(out)
}

// UsersList returns all users from DB
func UsersList(w http.ResponseWriter, r *http.Request) {
	repo, err := repo.NewUserRepo()
	if err != nil {
		panic(err)
	}

	items, err := repo.List()
	if err != nil {
		panic(err)
	}

	out, err := json.Marshal(items)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write(out)
}
