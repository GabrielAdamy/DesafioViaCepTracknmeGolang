package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"example.com/m/v2/adapters/inbound/dtos"
)

type ViaCepClient struct {
	//cep        dtos.ViaCepDTO
	httpClient *http.Client
}

func NewViaCepClient() *ViaCepClient {

	return &ViaCepClient{
		httpClient: &http.Client{},
	}
}

func (c *ViaCepClient) FindAddressByCep(cep string) (*dtos.ViaCepDTO, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%v/json/", cep)
	req, err := c.CreateRequest("GET", url)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var viaCep dtos.ViaCepDTO
	json.Unmarshal(bodyBytes, &viaCep)
	return &viaCep, nil
}

func (c *ViaCepClient) CreateRequest(method string, uri string) (*http.Request, error) {
	return http.NewRequest(method, uri, nil)
}
