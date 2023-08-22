package manager

import "github.com/NursiNursi/laundry-apps/repository"

type RepoManager interface {
	// semua repo di daftarkan disini
	UomRepo() repository.UomRepository
	ProductRepo() repository.ProductRepository
	CustomerRepo() repository.CustomerRepository
	EmployeeRepo() repository.EmployeeRepository
	BillRepo() repository.BillRepository
}

type repoManager struct {
	infra InfraManager
}

// BillRepo implements RepoManager.
func (r *repoManager) BillRepo() repository.BillRepository {
	return repository.NewBillRepository(r.infra.Conn())
}

// CustomerRepo implements RepoManager.
func (r *repoManager) CustomerRepo() repository.CustomerRepository {
	return repository.NewCustomerRepository(r.infra.Conn())
}

// EmployeeRepo implements RepoManager.
func (r *repoManager) EmployeeRepo() repository.EmployeeRepository {
	return repository.NewEmployeeRepository(r.infra.Conn())
}

// ProductRepo implements RepoManager.
func (r *repoManager) ProductRepo() repository.ProductRepository {
	return repository.NewProductRepository(r.infra.Conn())
}

// UomRepo implements RepoManager.
func (r *repoManager) UomRepo() repository.UomRepository {
	return repository.NewUomRepository(r.infra.Conn())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{infra: infra}
}
