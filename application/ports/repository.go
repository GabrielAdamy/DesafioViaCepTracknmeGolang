package ports

import "example.com/m/v2/application/domain"

type EmployeeRepository interface {
	FindAll(page int64, limit int64, parameters map[string]string) ([]*domain.Employee, int64, error)
	FindOne(id int64) (*domain.Employee, error)
	Save(employee *domain.Employee) error
	Delete(id int64) error
	Update(employee *domain.Employee) error

}