package repo

import (
	"database/sql"
	"fmt"

	models "github.com/glugate/uno/apps/crm/models"
	"github.com/glugate/uno/pkg/uno"
)

type Repo interface {
	// Create(o models.Menu) (*models.Menu, error)
	// Update(o models.Menu) (bool, error)
	// Delete(o models.Menu) (bool, error)
	// Has(id int64) bool
	Find(id string) (*models.Menu, error)
	List() ([]models.Menu, error)
}

type MenuRepo struct {
	stdDB *sql.DB
}

// NewMenuRepo grabs the database connection from already
// instatiated app and creates new Repo object with live db connection.
func NewMenuRepo() (Repo, error) {
	app := uno.Instance()
	if app == nil {
		return nil, fmt.Errorf("uno app not instatiated. Please run uno.NewNno()")
	}
	return &MenuRepo{
		stdDB: app.DB.StdDB,
	}, nil
}

// Find a model by its primary key
func (o *MenuRepo) Find(id string) (*models.Menu, error) {
	rows := o.stdDB.QueryRow("SELECT id, label from menus WHERE id=?", id)
	menu := models.Menu{}
	err := rows.Scan(
		&menu.ID,
		&menu.Label,
	)
	if err != nil {
		return nil, err
	}

	// Attach menu items for this menu
	menu.Items, err = o.ItemsList(menu.ID)
	if err != nil {
		return nil, err
	}

	return &menu, nil
}

// List retuns slice of all Menus from database
func (o *MenuRepo) List() (items []models.Menu, err error) {
	rows, err := o.stdDB.Query("SELECT id, label from menus")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		menu := models.Menu{}
		rows.Scan(
			&menu.ID,
			&menu.Label,
		)

		// Attach menu items for this menu
		menu.Items, err = o.ItemsList(menu.ID)
		if err != nil {
			return nil, err
		}
		items = append(items, menu)
	}
	return items, nil
}

/*

func (r *MenuRepo) Create(o models.Menu) (*models.Menu, error) {}

func (r *MenuRepo) Update(o models.Menu) (bool, error) {}

func (r *MenuRepo) Delete(o models.Menu) (bool, error) {}

func (r *MenuRepo) Has(id int64) bool {}

*/

// ItemsList retuns slice of all Menu Items for a menu from database
func (o *MenuRepo) ItemsList(id int64) (items []models.MenuItem, err error) {

	stmt, err := o.stdDB.Prepare("SELECT menu_id, label, path, ordering FROM menu_items WHERE menu_id = ?")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		item := models.MenuItem{}
		rows.Scan(
			&item.ID,
			&item.Label,
			&item.Path,
			&item.Ordering,
		)
		items = append(items, item)
	}
	return items, nil
}
