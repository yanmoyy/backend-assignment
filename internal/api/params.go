package api

type CreateIssueParams struct {
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	UserId      *uint  `json:"userId,omitempty"`
}
