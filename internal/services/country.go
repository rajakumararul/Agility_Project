package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FetchCountryByName(name string) (interface{}, error) {
	url := fmt.Sprintf("https://restcountries.com/v3.1/name/%s", name)
	return fetch(url)
}

func FetchCountryByCurrency(currency string) (interface{}, error) {
	url := fmt.Sprintf("https://restcountries.com/v3.1/currency/%s", currency)
	return fetch(url)
}

func FetchCountryByCapital(capital string) (interface{}, error) {
	url := fmt.Sprintf("https://restcountries.com/v3.1/capital/%s", capital)
	return fetch(url)
}

// Generic API fetcher
func fetch(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data, nil
}
