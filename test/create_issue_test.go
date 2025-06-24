package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yanmoyy/backend-assignment/internal/api"
	"github.com/yanmoyy/backend-assignment/internal/client"
)

func TestCreateIssue(t *testing.T) {
	cl := client.NewClient(baseURL)
	err := cl.Reset()
	if err != nil {
		t.Fatal(err)
	}
	t.Run("userId exists", func(t *testing.T) {
		id := uint(1)
		issue, err := cl.CreateIssue(
			api.CreateIssueParams{
				Title:       "버그 수정 필요",
				Description: "로그인 페이지에서 오류 발생",
				UserId:      &id,
			},
		)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, uint(1), issue.ID)
		assert.Equal(t, "버그 수정 필요", issue.Title)
		assert.Equal(t, "로그인 페이지에서 오류 발생", issue.Description)
		assert.Equal(t, api.StatusInProgress, issue.Status)
		assert.Equal(t, id, issue.User.ID)
		assert.Equal(t, "김개발", issue.User.Name)
	})
	t.Run("userId does not exist", func(t *testing.T) {
		issue, err := cl.CreateIssue(
			api.CreateIssueParams{
				Title:       "버그 수정 필요",
				Description: "로그인 페이지에서 오류 발생",
			},
		)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, uint(2), issue.ID)
		assert.Equal(t, api.StatusPending, issue.Status)
	})
}
