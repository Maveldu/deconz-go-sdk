package sdk

import (
	"fmt"
)

func (c Client) GetDevices() ([]string, error) {
	result := []string{}
	err := c.Get("/devices", nil, &result)
	if err  != nil {
		return nil, err
	}
	return result, nil
}

func (c Client) GetDevice(macAddress string) (Device, error) {
	result := Device{}
	err := c.Get(fmt.Sprintf("/devices/%s", macAddress), nil, &result)
	if err  != nil {
		return Device{}, err
	}
	return result, nil
}
