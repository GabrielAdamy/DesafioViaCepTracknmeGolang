package ports

import "example.com/m/v2/adapters/inbound/dtos"

type ViaCepClient interface {
	FindAddressByCep(cep string) (*dtos.ViaCepDTO, error)
}
