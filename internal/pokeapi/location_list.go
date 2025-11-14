package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (client *Client) ListAreaLocations(reqUrl *string) (AreaLocationsRes, error) {
	url := "https://pokeapi.co/api/v2/location-area"
	if reqUrl != nil {
		url = *reqUrl
	}

	if d, ok := client.cache.Get(url); ok {
		areaLocations := AreaLocationsRes{}
		err := json.Unmarshal(d, &areaLocations)
		if err != nil {
			return areaLocations, err
		}
		return areaLocations, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return AreaLocationsRes{}, err
	}

	res, err := client.httpClient.Do(req)
	if err != nil {
		return AreaLocationsRes{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return AreaLocationsRes{}, err
	}

	var areaLocations AreaLocationsRes
	err = json.Unmarshal(data, &areaLocations)
	if err != nil {
		return AreaLocationsRes{}, err
	}

	client.cache.Add(url, data)
	return areaLocations, nil
}
