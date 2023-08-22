package repository

import (
	"database/sql"

	"fmt"

	"github.com/NursiNursi/laundry-apps/model"
	"github.com/NursiNursi/laundry-apps/utils/exceptions"
)

type UomRepository interface {
	BaseRepository[model.Uom]
	GetByName(name string) (model.Uom, error)
}

type uomRepository struct {
	db *sql.DB
}

func (u *uomRepository) Create(payload model.Uom) error {
	_, err := u.db.Exec("INSERT INTO uom (id, name) VALUES ($1, $2)", payload.Id, payload.Name)
	exceptions.CheckErr(err)
	fmt.Println("UOM created successfully")
	return nil
}
func (u *uomRepository) List() ([]model.Uom, error) {
	rows, err := u.db.Query("SELECT id, name FROM uom")
	exceptions.CheckErr(err)

	var uoms []model.Uom
	for rows.Next() {
		var uom model.Uom
		err := rows.Scan(&uom.Id, &uom.Name)
		exceptions.CheckErr(err)
		uoms = append(uoms, uom)
	}
	fmt.Println("UOM successfully retreived")
	return uoms, nil
}

func (u *uomRepository) Get(id string) (model.Uom, error) {
	var uom model.Uom
	err := u.db.QueryRow("SELECT id, name FROM uom WHERE id=$1", id).Scan(&uom.Id, &uom.Name)
	exceptions.CheckErr(err)
	return uom, nil
}

func (u *uomRepository) GetByName(name string) (model.Uom, error) {
	var uom model.Uom
	// LIKE => case sensitive ILIKE => incase sensitive ()
	err := u.db.QueryRow("SELECT id, name FROM uom WHERE name ILIKE $1", "%"+name+"%").Scan(&uom.Id, &uom.Name)
	exceptions.CheckErr(err)
	return uom, nil
}

func (u *uomRepository) Update(payload model.Uom) error {
	_, err := u.db.Exec("UPDATE uom SET name=$1 WHERE id=$2", payload.Name, payload.Id)
	exceptions.CheckErr(err)
	return nil
}
func (u *uomRepository) Delete(id string) error {
	_, err := u.db.Exec("DELETE FROM uom WHERE id=$1", id)
	exceptions.CheckErr(err)
	return nil
}

// constructor
func NewUomRepository(db *sql.DB) UomRepository {
	return &uomRepository{db: db}
}