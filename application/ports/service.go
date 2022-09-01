package ports

import "example.com/m/v2/application/domain"

type EmployeeService interface {
	Save(dom *domain.Employee) error
	FindOne(id int64)  (*domain.Employee, error)
	FindAll(page, limit int64, parameters map[string]string) ([]*domain.Employee, int64 ,error)
	FindByCep(cep string, page, limit int64) ([]*domain.Employee, int64, error)
	Delete(id int64) error
	Update(dom *domain.Employee) error
	UpdatePatch(id int64, patch []byte) error
}