package commands

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type areaResponse struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(config *Config) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if config.Next != "" {
		url = config.Next
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var areas areaResponse

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&areas)
	if err != nil {
		return err
	}

	config.Next = areas.Next
	config.Previous = areas.Previous

	for _, r := range areas.Results {
		fmt.Println(r.Name)
	}

	return nil
}

func commandMapB(config *Config) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if config.Previous != "" {
		url = config.Previous
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var areas areaResponse

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&areas)
	if err != nil {
		return err
	}

	config.Next = areas.Next
	config.Previous = areas.Previous

	for _, r := range areas.Results {
		fmt.Println(r.Name)
	}

	return nil
}
