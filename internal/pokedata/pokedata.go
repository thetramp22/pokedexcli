package pokedata

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/thetramp22/pokedexcli/internal/pokecache"
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

func GetLocationArea(url string, cache *pokecache.Cache) (LocationArea, error) {
	data := []byte{}

	// get cached data if available
	if cachedData, ok := cache.Get(url); ok {
		data = cachedData
	} else {
		// fetch data
		res, err := http.Get(url)
		if err != nil {
			return LocationArea{}, err
		}
		fetchedData, err := io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			return LocationArea{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, fetchedData)
		}
		if err != nil {
			return LocationArea{}, err
		}
		data = fetchedData
	}

	// unmarshal data
	locationArea := LocationArea{}
	err := json.Unmarshal(data, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}
	cache.Add(url, data)
	return locationArea, nil
}
