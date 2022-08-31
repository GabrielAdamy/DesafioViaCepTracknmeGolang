package dtos

import "example.com/m/v2/application/domain"

type ViaCepDTO struct {
	Cep      string  `json:"cep"`
	Address  string `json:"logradouro"`
	District string `json:"bairro"`
	City     string `json:"localidade"`
	Uf       string `json:"uf"`
	Ibge     string `json:"ibge"`
	Ddd      string  `json:"ddd"`
}

func (c *ViaCepDTO) ParseToDomain() *domain.ViaCep {
	return &domain.ViaCep{
		Cep:      c.Cep,
		Address:  c.Address,
		District: c.District,
		City:     c.City,
		Uf:       c.Uf,
		Ibge:     c.Ibge,
		Ddd:      c.Ddd,
	}
}

type EmployeeDTO struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Age      int64  `json:"age"`
	Cep      string  `json:"cep"`
	Sex      string `json:"sex"`
	Address  string `json:"address"`
	District string `json:"district"`
	City     string `json:"city"`
	State    string `json:"state"`
}

func (c *EmployeeDTO) ParseToDomain() *domain.Employee {
	return &domain.Employee{
		ID:       c.ID,
		Name:     c.Name,
		Age:      c.Age,
		Sex:      c.Sex,
		Cep:      c.Cep,
		Address:  c.Address,
		District: c.District,
		City:     c.City,
		State:    c.State,
	}
}

func NewContractDTO(dom *domain.Employee) *EmployeeDTO {
	contract := &EmployeeDTO{
		ID:       dom.ID,
		Name:     dom.Name,
		Age:      dom.Age,
		Sex:      dom.Sex,
		Cep:      dom.Cep,
		Address:  dom.Address,
		District: dom.District,
		City:     dom.City,
		State:    dom.State,
	}
	return contract
}
