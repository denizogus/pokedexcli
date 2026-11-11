package main

import "fmt"

func commandMap(cfg *config) error {
	areas, err := cfg.client.GetLocationAreas(cfg.NextURL)
	if err != nil {
		return err
	}

	cfg.NextURL = areas.Next
	cfg.PrevURL = areas.Previous

	for _, location := range areas.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	areas, err := cfg.client.GetLocationAreas(cfg.PrevURL)
	if err != nil {
		return err
	}

	cfg.NextURL = areas.Next
	cfg.PrevURL = areas.Previous

	for _, location := range areas.Results {
		fmt.Println(location.Name)
	}

	return nil
}
