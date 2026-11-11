package main

import "fmt"

func commandMap(cfg *config, param []string) error {
	if len(param) != 0 {
		fmt.Errorf("wrong parameters")
	}
	areas, err := cfg.client.GetLocationAreas20(cfg.NextURL)
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

func commandMapb(cfg *config, param []string) error {
	areas, err := cfg.client.GetLocationAreas20(cfg.PrevURL)
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
