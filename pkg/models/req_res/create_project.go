package req_res

type CreateProjectRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Lead        string `json:"lead" binding:"required"`
}
