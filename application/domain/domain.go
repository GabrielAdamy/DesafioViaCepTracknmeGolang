package domain

type ViaCep struct {
	Cep      string `json:"cep"`
	Address  string `json:"logradouro"`
	District string `json:"bairro"`
	City     string `json:"localidade"`
	Uf       string `json:"uf"`
	Ibge     string `json:"ibge"`
	Ddd      string  `json:"ddd"`
}

type Employee struct {
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
