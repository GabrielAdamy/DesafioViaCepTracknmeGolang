package services

import (
	"encoding/json"
	"time"

	"example.com/m/v2/application/domain"
	"example.com/m/v2/application/ports"
	jsonpatch "github.com/evanphx/json-patch/v5"
)

type EmployeeServiceImpl struct {
	repository ports.EmployeeRepository
	client     ports.ViaCepClient
}

func NewEmployeeServiceImpl(repository ports.EmployeeRepository, client ports.ViaCepClient) *EmployeeServiceImpl {
	return &EmployeeServiceImpl{
		repository: repository,
		client:     client,
	}
}

func (c *EmployeeServiceImpl) Save(dom *domain.Employee) error {
	dom.ID = time.Now().UnixMilli()
	c.UpdateCep(dom.Cep, dom)
	return c.repository.Save(dom)
}

func (c *EmployeeServiceImpl) FindOne(id int64) (*domain.Employee, error) {
	return c.repository.FindOne(id)
}

func (c *EmployeeServiceImpl) FindByCep(cep string, page, limit int64) ([]*domain.Employee, int64, error) {
	return c.repository.FindByCep(cep, page, limit)
}

func (c *EmployeeServiceImpl) FindAll(page, limit int64, parameters map[string]string) ([]*domain.Employee, int64, error) {
	return c.repository.FindAll(page, limit, parameters)
}

func (c *EmployeeServiceImpl) Delete(id int64) error {
	return c.repository.Delete(id)
}

func (c *EmployeeServiceImpl) Update(dom *domain.Employee) error {
	return c.repository.Update(dom)
}

func (c *EmployeeServiceImpl) UpdatePatch(id int64, patch []byte) error {
	employee, err := c.FindOne(id)
	if err != nil {
		return err
	}

	bytes, err := json.Marshal(employee)
	if err != nil {
		return err
	}

	modified, err := jsonpatch.MergePatch(bytes, patch)
	if err != nil {
		return err
	}

	json.Unmarshal(modified, &employee)
	c.UpdateCep(employee.Cep, employee)
	err = c.repository.Update(employee)
	if err != nil {
		return err
	}

	return nil
}

func (c *EmployeeServiceImpl) UpdateCep(cep string, dom *domain.Employee) error {
	address, err := c.client.FindAddressByCep(cep)
	if err != nil {
		return err
	}
	dom.Address = address.Address
	dom.District = address.District
	dom.Cep = cep
	dom.State = address.Uf
	dom.City = address.City
	return nil
}
