package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
)

type API string

const (
	CDN              API    = "CDN"
	VIA              API    = "VIA"
	INVALID_ENDPOINT string = "INVALID_ENDPOINT"
)

type CepInput struct {
	Cep string
	Api API
}

func (c *CepInput) GetCepRequest() (*http.Request, error) {
	var url string
	switch c.Api {
	case CDN:
		url = "https://cdn.apicep.com/file/apicep/" + c.Cep + ".json"
	case VIA:
		url = "http://viacep.com.br/ws/" + c.Cep + "/json/"
	default:
		return nil, errors.New(INVALID_ENDPOINT)
	}

	return http.NewRequest(http.MethodGet, url, nil)
}

func ValidateCep(cep string) (string, error) {
	isValidCPF, err := regexp.Match("^\\d{5}-\\d{3}$", []byte(cep))

	if err != nil {
		return "", err
	}

	if !isValidCPF {
		return "", fmt.Errorf("%s is not a valid CPF", cep)
	}
	return cep, nil
}
