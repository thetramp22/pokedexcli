package pokeData

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationArea struct {
	Count    int       `json:"count"`
	Next     *string   `json:"next"`
	Previous *string   `json:"previous"`
	Results  []Results `json:"results"`
}

type Results struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func GetLocationArea(url string) (LocationArea, error) {
	res, err := http.Get(url)
	if err != nil {
		return LocationArea{}, err
	}
	data, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return LocationArea{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, data)
	}
	if err != nil {
		return LocationArea{}, err
	}
	locationArea := LocationArea{}
	err = json.Unmarshal(data, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}
	return locationArea, nil
}
