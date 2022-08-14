package repo

import (
	"database/sql"
	"fmt"

	models "github.com/glugate/uno/apps/crm/models"
	"github.com/glugate/uno/pkg/uno"
)

type IUsersRepo interface {
	Find(id string) (*models.User, error)
	List() ([]models.User, error)
}

type UsersRepo struct {
	stdDB *sql.DB
}

// NewUserRepo grabs the database connection from already
// instatiated app and creates new Repo object with live db connection.
func NewUserRepo() (IUsersRepo, error) {
	app := uno.Instance()
	if app == nil {
		return nil, fmt.Errorf("uno app not instatiated. Please run uno.NewNno()")
	}
	return &UsersRepo{
		stdDB: app.DB.StdDB,
	}, nil
}

// Find a model by its primary key
func (o *UsersRepo) Find(id string) (*models.User, error) {
	stmt, err := o.stdDB.Prepare("SELECT id, first_name, last_name, email FROM users WHERE id = ?")
	if err != nil {
		return nil, err
	}
	rows := stmt.QueryRow(id)

	// Scan into new model
	user := models.User{}
	err = rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// List retuns slice of all Users from database
func (o *UsersRepo) List() (items []models.User, err error) {
	stmt, err := o.stdDB.Prepare("SELECT id, first_name, last_name, email, birth_date, title, status, is_active FROM users")
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		user := models.User{}
		rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.BirthDate,
			&user.Title,
			&user.Status,
			&user.IsActive,
		)
		items = append(items, user)
	}
	return items, nil
}
