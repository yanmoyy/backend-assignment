package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yanmoyy/backend-assignment/internal/api"
	"github.com/yanmoyy/backend-assignment/internal/db"
)

func Test_updateIssue(t *testing.T) {
	t.Run("update title description", func(t *testing.T) {
		issue := api.Issue{
			ID:          1,
			Title:       "title",
			Description: "description",
			Status:      api.StatusPending,
			User:        nil,
		}
		updatedIssue, err := updateIssue(
			api.UpdateIssueParams{
				Title:       "new title",
				Description: "new description",
			},
			issue,
		)
		assert.NoError(t, err)
		assert.Equal(t, "new title", updatedIssue.Title)
		assert.Equal(t, "new description", updatedIssue.Description)
	})
	t.Run("need to specify userId", func(t *testing.T) {
		issue := api.Issue{
			ID:          1,
			Title:       "title",
			Description: "description",
			Status:      api.StatusPending,
		}
		_, err := updateIssue(
			api.UpdateIssueParams{
				Status: api.StatusInProgress,
			},
			issue,
		)
		assert.Error(t, err)
	})
	t.Run("status: PENDING", func(t *testing.T) {
		issue := api.Issue{
			ID:          1,
			Title:       "title",
			Description: "description",
			Status:      api.StatusPending,
			User:        nil,
		}
		id := uint(1)
		updated, err := updateIssue(
			api.UpdateIssueParams{
				UserId: &id,
			},
			issue,
		)
		assert.NoError(t, err)
		assert.Equal(t, api.StatusInProgress, updated.Status)
	})
	t.Run("(userId -> nil)", func(t *testing.T) {
		issue := api.Issue{
			ID:          1,
			Title:       "title",
			Description: "description",
			Status:      api.StatusInProgress,
			User:        &api.User{},
		}
		updated, err := updateIssue(
			api.UpdateIssueParams{
				UserId: nil,
			},
			issue,
		)
		assert.NoError(t, err)
		assert.Equal(t, api.StatusPending, updated.Status)
	})
	t.Run("status: COMPLETED or CANCELLED", func(t *testing.T) {
		issue := api.Issue{
			Status: api.StatusCompleted,
		}
		_, err := updateIssue(api.UpdateIssueParams{Title: "new title"}, issue)
		assert.Error(t, err)

		issue = api.Issue{
			Status: api.StatusCancelled,
		}
		_, err = updateIssue(api.UpdateIssueParams{Title: "new title"}, issue)
		assert.Error(t, err)
	})
	t.Run("set user", func(t *testing.T) {
		issue := api.Issue{
			ID:     1,
			Status: api.StatusPending,
			User:   nil,
		}
		id := uint(1)
		updated, err := updateIssue(
			api.UpdateIssueParams{
				UserId: &id,
			},
			issue,
		)
		assert.NoError(t, err)
		assert.Equal(t, api.StatusInProgress, updated.Status)
		assert.Equal(t, id, updated.User.ID)
		assert.Equal(t, db.GetUserByID(id).Name, updated.User.Name)
	})
}
