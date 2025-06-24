package client

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/yanmoyy/backend-assignment/internal/api"
)

type CreateIssueParams struct {
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	UserId      *uint  `json:"userId,omitempty"`
}

func (c *Client) CreateIssue(body CreateIssueParams) (api.Issue, error) {
	url := c.baseURL + "/issue"
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return api.Issue{}, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(bodyBytes))
	if err != nil {
		return api.Issue{}, err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return api.Issue{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 201 {
		return api.Issue{}, err
	}
	var issue api.Issue
	err = json.NewDecoder(resp.Body).Decode(&issue)
	if err != nil {
		return api.Issue{}, err
	}
	return issue, nil
}
