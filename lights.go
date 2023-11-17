package sdk

import (
	"fmt"
)

func (c Client) GetLights() (map[string]Light, error) {
	result := map[string]Light{}
	err := c.Get("/lights", nil, &result)
	if err  != nil {
		return nil, err
	}
	return result, nil
}

func (c Client) GetLight(id string) (Light, error) {
	result := Light{}
	err := c.Get(fmt.Sprintf("/lights/%s", id), nil, &result)
	if err  != nil {
		return Light{}, err
	}
	return result, nil
}
