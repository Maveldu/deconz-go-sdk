package sdk

func (c Client) GetConfig() (Config, error) {
	result := Config{}
	err := c.Get("/config", nil, &result)
	if err  != nil {
		return Config{}, err
	}
	return result, nil
}
