package global

//用户分页条件
type UserQuery struct {
	PageNum      int
	PageSize     int
	Keywords     string
	OrganizatiId int
}

//组织机构
type Organization struct {
	BaseModel
	Name string `gorm:"comment:名称" json:"name"`
}
