package sdk

import (
	"fmt"
)

func (c Client) GetSensors() (map[string]Sensor, error) {
	result := map[string]Sensor{}
	err := c.Get("/sensors", nil, &result)
	if err  != nil {
		return nil, err
	}
	return result, nil
}

func (c Client) GetSensor(id string) (Sensor, error) {
	result := Sensor{}
	err := c.Get(fmt.Sprintf("/sensors/%s", id), nil, &result)
	if err  != nil {
		return Sensor{}, err
	}
	return result, nil
}
