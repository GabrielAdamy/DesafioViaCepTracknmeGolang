package repositories

import (
	"fmt"

	"example.com/m/v2/adapters/configurations"
	"example.com/m/v2/application/domain"
)

type EmployeeRepositoryImpl struct {
	db *configurations.PostgresConfiguration
}

func NewEmployeeRepositoryImpl(db *configurations.PostgresConfiguration) *EmployeeRepositoryImpl {
	db.Migrate(&domain.Employee{})

	return &EmployeeRepositoryImpl{db: db}
}

// retorna o slice de registros, count e erro
func (rep *EmployeeRepositoryImpl) FindAll(page int64, limit int64, parameters map[string]string) ([]*domain.Employee, int64, error) {
	query := "1=1"
	var values []interface{}
	if len(parameters) > 0 {
		if parameters["name"] != "" {
			if parameters["name"] == "null" {
				query += " AND name is null"
			} else {
				query += " AND name = ?"
				values = append(values, parameters["name"])
			}
		}

		if parameters["cep"] != "" {
			if parameters["cep"] == "null" {
				query += " AND cep is null"
			} else {
				query += " AND cep = ?"
				values = append(values, parameters["cep"])
			}
		}
		if parameters["id"] != "" {
			query += " AND id = ?"
			values = append(values, parameters["id"])
		}
	}

	if page <= 0 {
		page = 1
	}

	var entities []*domain.Employee
	tx := rep.db.DB.Where(query, values...).Order("id desc").Limit(int(limit)).Offset(int((page - 1) * limit)).Find(&entities)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}

	var totalRows int64
	rep.db.DB.Model(&domain.Employee{}).Where(query, values...).Count(&totalRows)

	return entities, totalRows, nil
}

func (rep *EmployeeRepositoryImpl) FindOne(id int64) (*domain.Employee, error) {
	var entity domain.Employee
	if tx := rep.db.DB.First(&entity, id); tx.Error != nil {
		return nil, tx.Error
	}
	return &entity, nil
}

func (rep *EmployeeRepositoryImpl) Save(employee *domain.Employee) error {
	tx := rep.db.DB.Create(employee)
	return tx.Error
}

func (rep *EmployeeRepositoryImpl) Delete(id int64) error {
	delete := fmt.Sprintf("DELETE FROM employees WHERE ID=%v", id)
	tx := rep.db.DB.Exec(delete)

	return tx.Error
}

func (rep *EmployeeRepositoryImpl) Update(employee *domain.Employee) error {
	tx := rep.db.DB.Save(employee)
	return tx.Error
}


