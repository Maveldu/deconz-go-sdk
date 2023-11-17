package sdk

import (
	"fmt"
)

func (c Client) GetRules() (map[string]Rule, error) {
	result := map[string]Rule{}
	err := c.Get("/rules", nil, &result)
	if err  != nil {
		return nil, err
	}
	return result, nil
}

func (c Client) GetRule(id string) (Rule, error) {
	result := Rule{}
	err := c.Get(fmt.Sprintf("/rules/%s", id), nil, &result)
	if err  != nil {
		return Rule{}, err
	}
	return result, nil
}

