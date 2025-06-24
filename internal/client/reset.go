package client

import (
	"fmt"
	"net/http"
)

func (c *Client) Reset() error {
	url := c.baseURL + "/reset"
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return fmt.Errorf("create request failed")
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code is not 200")
	}
	return nil
}
