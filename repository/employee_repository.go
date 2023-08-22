package repository

import (
	"database/sql"

	"github.com/NursiNursi/laundry-apps/model"
	"github.com/NursiNursi/laundry-apps/model/dto"
	"github.com/NursiNursi/laundry-apps/utils/common"
)

type EmployeeRepository interface {
	BaseRepository[model.Employee]
	BaseRepositoryPaging[model.Employee]
	GetPhoneNumber(phoneNumber string) (model.Employee, error)
}

type employeeRepository struct {
	db *sql.DB
}

// Create implements employeeRepository.
func (e *employeeRepository) Create(payload model.Employee) error {
	_, err := e.db.Exec("INSERT INTO employee (id, name, phone_number, address) VALUES ($1, $2, $3, $4)", payload.Id, payload.Name, payload.PhoneNumber, payload.Address)
	if err != nil {
		return err
	}
	return nil
}

// Delete implements employeeRepository.
func (e *employeeRepository) Delete(id string) error {
	_, err := e.db.Exec("DELETE FROM employee WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

// Get implements employeeRepository.
func (e *employeeRepository) Get(id string) (model.Employee, error) {
	var employee model.Employee
	err := e.db.QueryRow("SELECT id, name, phone_number, address FROM employee WHERE id=$1", id).Scan(&employee.Id, &employee.Name, &employee.PhoneNumber, &employee.Address)
	if err != nil {
		return model.Employee{}, err
	}
	return employee, nil

}

// GetEmail implements employeeRepository.
func (e *employeeRepository) GetPhoneNumber(phoneNumber string) (model.Employee, error) {
	var employee model.Employee
	err := e.db.QueryRow("SELECT id, name, phone_number, address FROM employee WHERE phone_number=$1", phoneNumber).Scan(&employee.Id, &employee.Name, &employee.PhoneNumber, &employee.Address)
	if err != nil {
		return model.Employee{}, err
	}
	return employee, nil
}

// List implements employeeRepository.
func (e *employeeRepository) List() ([]model.Employee, error) {
	rows, err := e.db.Query("SELECT id, name, phone_number, address FROM employee")
	if err != nil {
		return nil, err
	}
	var employees []model.Employee
	for rows.Next() {
		var employee model.Employee
		err := rows.Scan(&employee.Id, &employee.Name, &employee.PhoneNumber, &employee.Address)
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}
	return employees, nil
}

// Paging implements employeeRepository.
func (e *employeeRepository) Paging(requestPaging dto.PaginationParam) ([]model.Employee, dto.Paging, error) {
	var paginationQuery dto.PaginationQuery
	paginationQuery = common.GetPaginationParams(requestPaging)
	rows, err := e.db.Query("SELECT id, name, phone_number, address FROM employee LIMIT $1 OFFSET $2", paginationQuery.Take, paginationQuery.Skip)
	if err != nil {
		return nil, dto.Paging{}, err
	}
	var employees []model.Employee
	for rows.Next() {
		var employee model.Employee
		err := rows.Scan(&employee.Id, &employee.Name, &employee.PhoneNumber, &employee.Address)
		if err != nil {
			return nil, dto.Paging{}, err
		}
		employees = append(employees, employee)
	}

	// count product
	var totalRows int
	row := e.db.QueryRow("SELECT COUNT(*) FROM employee")
	err = row.Scan(&totalRows)
	if err != nil {
		return nil, dto.Paging{}, err
	}

	return employees, common.Paginate(paginationQuery.Page, paginationQuery.Take, totalRows), nil
}

// Update implements employeeRepository.
func (e *employeeRepository) Update(payload model.Employee) error {
	_, err := e.db.Exec("UPDATE employee SET name = $2, phone_number = $3, address = $4 WHERE id = $1", payload.Id, payload.Name, payload.PhoneNumber, payload.Address)
	if err != nil {
		return err
	}
	return nil
}

func NewEmployeeRepository(db *sql.DB) EmployeeRepository {
	return &employeeRepository{db: db}
}
