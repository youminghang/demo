package forms

type AddTagForm struct {
	Name      string `form:"name" json:"name" binding:"required,min=1,max=100"`
	CreatedBy string `form:"created_by" json:"created_by" binding:"required,min=1,max=100"`
	State     *int   `form:"state" json:"state" binding:"required,oneof=0 1"`
}

type UpdateTagForm struct {
	Name      string `form:"name" json:"name" binding:"required,min=1,max=100"`
	UpdatedBy string `form:"updated_by" json:"updated_by" binding:"required,min=1,max=100"`
	State     *int   `form:"state" json:"state" binding:"omitempty,oneof=0 1"`
}
