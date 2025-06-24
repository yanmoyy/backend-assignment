package client

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/yanmoyy/backend-assignment/internal/api"
)

func (c *Client) CreateIssue(body api.CreateIssueParams) (api.Issue, error) {
	url := c.baseURL + "/issue"
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return api.Issue{}, err
	}
	resp, err := c.client.Post(url, "application/json", bytes.NewReader(bodyBytes))
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

func (c *Client) GetIssue(id int) (api.Issue, error) {
	url := c.baseURL + "/issue/" + strconv.Itoa(id)
	resp, err := c.client.Get(url)
	if err != nil {
		return api.Issue{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return api.Issue{}, err
	}
	var issue api.Issue
	err = json.NewDecoder(resp.Body).Decode(&issue)
	if err != nil {
		return api.Issue{}, err
	}
	return issue, nil
}

func (c *Client) UpdateIssue(id int, body api.UpdateIssueParams) (api.Issue, error) {
	url := c.baseURL + "/issue/" + strconv.Itoa(id)
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return api.Issue{}, err
	}
	req, err := http.NewRequest("PATCH", url, bytes.NewReader(bodyBytes))
	if err != nil {
		return api.Issue{}, err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return api.Issue{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return api.Issue{}, err
	}
	var issue api.Issue
	err = json.NewDecoder(resp.Body).Decode(&issue)
	if err != nil {
		return api.Issue{}, err
	}
	return issue, nil
}
