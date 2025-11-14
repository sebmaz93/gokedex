package repl

import (
	"errors"
	"fmt"
)

func commandMap(config *Config, param1 string) error {
	areaLocationsRes, err := config.ApiClient.ListAreaLocations(config.NextUrl)
	if err != nil {
		return err
	}

	config.NextUrl = areaLocationsRes.Next
	config.PreviousUrl = areaLocationsRes.Previous

	for _, r := range areaLocationsRes.Results {
		fmt.Println(r.Name)
	}
	fmt.Println("")

	return nil
}

func commandMapB(config *Config, param1 string) error {
	if config.PreviousUrl == nil {
		return errors.New("you're on the first page")
	}
	areaLocationsRes, err := config.ApiClient.ListAreaLocations(config.PreviousUrl)
	if err != nil {
		return err
	}

	config.NextUrl = areaLocationsRes.Next
	config.PreviousUrl = areaLocationsRes.Previous

	for _, r := range areaLocationsRes.Results {
		fmt.Println(r.Name)
	}
	fmt.Println("")

	return nil
}
