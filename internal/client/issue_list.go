package client

import (
	"encoding/json"
	"net/http"

	"github.com/yanmoyy/backend-assignment/internal/api"
)

func (c *Client) GetIssuesList(status api.IssueStatus) ([]api.Issue, error) {
	url := c.baseURL + "/issues"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	if status != "" {
		q := req.URL.Query()
		q.Add("status", string(status))
		req.URL.RawQuery = q.Encode()
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}
	var response api.GetIssuesResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return response.Issues, nil
}
