package sdk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/go-querystring/query"
)

type Error struct {
	Status     string
	StatusCode int
}

func (e Error) Error() string {
	return e.Status
}

type Client struct {
	http.Client

	url string
}

func NewClient(url string) *Client {
	return &Client{url: url}
}

func (c Client) Get(path string, params interface{}, result interface{}) error {
	paramsStr := ""
	if params != nil {
		p, err := query.Values(params)
		if err != nil {
			panic(fmt.Sprintf("query.Values failed with params %+v", params))
		}
		paramsStr = "?" + p.Encode()
	}

	resp, err := c.Client.Get(fmt.Sprintf("%s%s%s", c.url, path, paramsStr))
	if err != nil {
		return err
	}

	return handleResp(resp, result)
}

func handleResp(resp *http.Response, result interface{}) error {
	if resp.StatusCode >= 300 {
		return &Error{Status: resp.Status, StatusCode: resp.StatusCode}
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, result)
	if err != nil {
		return fmt.Errorf("could not decode body %q: %w", string(body), err)
	}

	return nil
}
