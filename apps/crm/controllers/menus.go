package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/glugate/uno/apps/crm/repo"
	"github.com/glugate/uno/pkg/uno/server"
)

// MenusFind returns one menu ( not menu item as child of menu )
// e.g. Admin Menu, Settings Menu
func MenusFind(w http.ResponseWriter, r *http.Request) {
	id := server.RequestParam(r, 0)
	repo, err := repo.NewMenuRepo()
	if err != nil {
		panic(err)
	}
	items, err := repo.Find(id)
	if err != nil {
		panic(err)
	}

	// Get the menu items

	out, err := json.Marshal(items)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write(out)
}

// MenusList returns all menus from DB
func MenusList(w http.ResponseWriter, r *http.Request) {
	repo, err := repo.NewMenuRepo()
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
