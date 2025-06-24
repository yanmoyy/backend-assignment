package client

import (
	"fmt"
	"net/http"
)

func (c *Client) Reset() error {
	url := c.baseURL + "/reset"
	resp, err := http.DefaultClient.Post(url, "application/json", nil)
	if err != nil {
		return fmt.Errorf("request failed")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code is not 200")
	}
	return nil
}
