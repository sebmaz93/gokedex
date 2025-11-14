package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (client *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name
	if d, ok := client.cache.Get(url); ok {
		pokemon := Pokemon{}
		err := json.Unmarshal(d, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, nil
	}

	res, err := client.httpClient.Do(req)
	defer res.Body.Close()
	if err != nil {
		return Pokemon{}, err
	}

	pokemen := Pokemon{}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}
	err = json.Unmarshal(data, &pokemen)
	if err != nil {
		return Pokemon{}, err
	}

	client.cache.Add(url, data)
	return pokemen, nil
}
