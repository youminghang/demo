package forms

type AddArticleForm struct {
	TagId     int32  `form:"tag_id" json:"tag_id" binding:"required,min=1"`
	Title     string `form:"title" json:"title" binding:"required,min=1,max=100"`
	Desc      string `form:"desc" json:"desc" binding:"required,min=1,max=255"`
	Content   string `form:"content" json:"content" binding:"required,min=1"`
	CreatedBy string `form:"created_by" json:"created_by" binding:"required,min=1,max=100"`
	State     *int   `form:"state" json:"state" binding:"required,oneof=0 1"`
}

type EditeArticleForm struct {
	TagId     *int32  `form:"tag_id" json:"tag_id" binding:"omitempty,min=1"`
	Title     string `form:"title" json:"title" binding:"omitempty,min=1,max=100"`
	Desc      string `form:"desc" json:"desc" binding:"omitempty,min=1,max=255"`
	Content   string `form:"content" json:"content" binding:"omitempty,min=1"`
	UpdatedBy string `form:"updated_by" json:"updated_by" binding:"required,min=1,max=100"`
	State     *int   `form:"state" json:"state" binding:"omitempty,oneof=0 1"`
}
