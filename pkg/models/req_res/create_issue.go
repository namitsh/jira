package req_res

type CreateIssueRequest struct {
	Description string `json:"description" binding:"required""`
	Summary     string `json:"summary" binding:"required"`
	Type        string `json:"type" binding:"required"`
}
