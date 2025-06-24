package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yanmoyy/backend-assignment/internal/api"
	"github.com/yanmoyy/backend-assignment/internal/client"
)

func TestGetIssues(t *testing.T) {
	cl := client.NewClient(baseURL)
	setupIssues(t, cl)
	t.Run("no status", func(t *testing.T) {
		issues, err := cl.GetIssuesList("")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, 2, len(issues))
	})
	t.Run("status = IN_PROGRESS", func(t *testing.T) {
		issues, err := cl.GetIssuesList(api.StatusInProgress)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, 1, len(issues))
		assert.Equal(t, "title1", issues[0].Title)
		assert.Equal(t, "description1", issues[0].Description)
	})
	t.Run("status = PENDING", func(t *testing.T) {
		issues, err := cl.GetIssuesList(api.StatusPending)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, 1, len(issues))
		assert.Equal(t, "title2", issues[0].Title)
		assert.Equal(t, "description2", issues[0].Description)
	})
}
