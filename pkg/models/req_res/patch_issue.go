package req_res

type PatchIssueRequest struct {
	Summary     *string `json:"summary,omitempty" bson:"summary,omitempty" binding:"omitempty"`
	Description *string `json:"description,omitempty" bson:"description,omitempty" binding:"omitempty"`
	Status      *string `json:"status,omitempty" bson:"status,omitempty" binding:"omitempty"`
	//UpdatedAt   time.Time `bson:"updated_at" json:"-"`
}
