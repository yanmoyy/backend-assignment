package api

type CreateIssueParams struct {
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	UserId      *uint  `json:"userId,omitempty"`
}

type UpdateIssueParams struct {
	Title       string      `json:"title,omitempty"`
	Description string      `json:"description,omitempty"`
	Status      IssueStatus `json:"status,omitempty"`
	UserId      *uint       `json:"userId,omitempty"`
}
