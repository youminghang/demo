package models

// 标签表
type Tag struct {
	BaseModel
	Name      string `gorm:"type:varchar(100);not null" json:"name"`       // 标签名称
	CreatedBy string `gorm:"type:varchar(100);not null" json:"created_by"` // 创建人
	UpdatedBy string `gorm:"type:varchar(100);not null" json:"updated_by"` // 修改人
}

// 文章表
type Article struct {
	BaseModel
	TagId     int32 `gorm:"type:int;not null"`
	Tag       Tag
	Title     string `gorm:"type:varchar(100);not null" json:"title"` // 文章标题
	Desc      string `gorm:"type:varchar(255);not null" json:"desc"`  // 简述
	Content   string `gorm:"type:text;not null" json:"content"`
	CreatedBy string `gorm:"type:varchar(100);not null" json:"created_by"` // 创建人
	UpdatedBy string `gorm:"type:varchar(100);not null" json:"updated_by"` // 修改人

}

// 认证表
type Auth struct {
	BaseModel
	UserName string `gorm:"type:varchar(100);not null" json:"user_name"` // 用户名
	PassWord string `gorm:"type:varchar(100);not null" json:"pass_word"` // 密码
}
