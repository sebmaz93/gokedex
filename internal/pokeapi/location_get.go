package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (client *Client) ExploreLocation(name string) (ExploreLocationRes, error) {
	url := baseURL + "/location-area/" + name
	if d, ok := client.cache.Get(url); ok {
		exploreLocation := ExploreLocationRes{}
		err := json.Unmarshal(d, &exploreLocation)
		if err != nil {
			return ExploreLocationRes{}, err
		}
		return exploreLocation, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ExploreLocationRes{}, nil
	}

	res, err := client.httpClient.Do(req)
	defer res.Body.Close()

	exploreLocation := ExploreLocationRes{}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return ExploreLocationRes{}, err
	}
	err = json.Unmarshal(data, &exploreLocation)
	if err != nil {
		return ExploreLocationRes{}, err
	}
	client.cache.Add(url, data)

	return exploreLocation, nil
}
