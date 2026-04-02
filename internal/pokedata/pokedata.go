package pokedata

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/thetramp22/pokedexcli/internal/pokecache"
)

type LocationAreas struct {
	Count    int       `json:"count"`
	Next     *string   `json:"next"`
	Previous *string   `json:"previous"`
	Results  []Results `json:"results"`
}

type Results struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func GetLocationAreas(url string, cache *pokecache.Cache) (LocationAreas, error) {
	data := []byte{}

	// get cached data if available
	if cachedData, ok := cache.Get(url); ok {
		data = cachedData
	} else {
		// fetch data
		res, err := http.Get(url)
		if err != nil {
			return LocationAreas{}, err
		}
		fetchedData, err := io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			return LocationAreas{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, fetchedData)
		}
		if err != nil {
			return LocationAreas{}, err
		}
		data = fetchedData
	}

	// unmarshal data
	locationAreas := LocationAreas{}
	err := json.Unmarshal(data, &locationAreas)
	if err != nil {
		return LocationAreas{}, err
	}
	cache.Add(url, data)
	return locationAreas, nil
}

type LocationArea struct {
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
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

type Pokemon struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
}

func GetPokemon(url string, cache *pokecache.Cache) (Pokemon, error) {
	data := []byte{}

	// get cached data if available
	if cachedData, ok := cache.Get(url); ok {
		data = cachedData
	} else {
		// fetch data
		res, err := http.Get(url)
		if err != nil {
			return Pokemon{}, err
		}
		fetchedData, err := io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			return Pokemon{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, fetchedData)
		}
		if err != nil {
			return Pokemon{}, err
		}
		data = fetchedData
	}

	// unmarshal data
	pokemon := Pokemon{}
	err := json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}
	cache.Add(url, data)
	return pokemon, nil
}
