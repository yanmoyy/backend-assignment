package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const baseURL = "http://localhost:8080"

func TestCreateIssue(t *testing.T) {
	url := baseURL + "/issue"
	id := uint(1)
	body := CreateIssueParmas{
		Title:       "버그 수정 필요",        // 필수
		Description: "로그인 페이지에서 오류 발생", // 선택
		UserId:      &id,
	}
	bodyData, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyData))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 201 {
		t.Fatalf("Expected status code 201, got %d", resp.StatusCode)
	}
	var issue Issue
	err = json.NewDecoder(resp.Body).Decode(&issue)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(issue)
	assert.Equal(t, body.Title, issue.Title)
	assert.Equal(t, body.Description, issue.Description)
	assert.Equal(t, StatusInProgress, issue.Status)
	assert.Equal(t, id, issue.User.ID)
}

func TestGetIssuesList(t *testing.T) {
	url := baseURL + "/issues"
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %d", resp.StatusCode)
	}
	var response struct {
		Issues []Issue `json:"issues"`
	}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1, len(response.Issues))
}
