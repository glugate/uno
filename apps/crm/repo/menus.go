package repo

import (
	"database/sql"
	"fmt"

	models "github.com/glugate/uno/apps/crm/models"
	"github.com/glugate/uno/pkg/uno/db"
	"gorm.io/gorm"
)

type Repo interface {
	Create(o models.Menu) (*models.Menu, error)
	Update(o models.Menu) (bool, error)
	Delete(o models.Menu) (bool, error)
	Has(id int64) bool
	Find(id int64) (*models.Menu, error)
	List() ([]*models.Menu, error)
}

type MenuRepo struct {
	DB      *sql.DB
	adapter *db.Adapter
}

func NewMenuRepo(stdDB *sql.DB, dsn string) Repo {
	return &MenuRepo{
		DB:      stdDB,
		adapter: db.NewAdapter(dsn, stdDB),
	}
}

func (r *MenuRepo) Create(o models.Menu) (*models.Menu, error) {
	err := r.adapter.Create(&o).Error
	if err != nil {
		return &o, err
	}
	return &o, nil
}

func (r *MenuRepo) Where(w any) *gorm.DB {
	return r.adapter.Where(w)
}

func (r *MenuRepo) Update(o models.Menu) (bool, error) {
	originMenu := new(models.Menu)

	err := r.adapter.Where("id=?", o.ID).First(originMenu).Updates(map[string]interface{}{
		"label": o.Label + " - updated",
	}).Error

	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *MenuRepo) Delete(o models.Menu) (bool, error) {
	fmt.Println(o.ID)
	err := r.adapter.Where("id=?", o.ID).Delete(&models.Menu{}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *MenuRepo) Has(id int64) bool {
	err := r.adapter.Where("id=?", id).First(models.Menu{}).Error
	return err == nil
}

func (r *MenuRepo) Find(id int64) (*models.Menu, error) {
	product := new(models.Menu)

	err := r.adapter.Where("id=?", id).First(product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *MenuRepo) List() (products []*models.Menu, err error) {
	err = r.adapter.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}
