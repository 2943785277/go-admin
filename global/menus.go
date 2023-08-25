package global

//路由
type Routes struct {
	Menus
	Children []Routes                          `gorm:"comment:路由" json:"children"` //
	Meta     `gorm:"comment:路由地址" json:"meta"` //
}

type Meta struct {
	Title     string `gorm:"comment:路由地址" json:"title"`
	Icon      string `gorm:"comment:路由地址" json:"icon"`
	Hidden    bool   `gorm:"comment:路由地址" json:"hidden"` //是否隐藏，不显示
	Roles     string `gorm:"comment:路由地址" json:"roles"`
	KeepAlive bool   `gorm:"comment:路由地址" json:"KeepAlive"`
}

//路由
type Menus struct {
	BaseModel
	Path     string `gorm:"comment:路由地址" json:"path"`       //路由地址
	Redirect string `gorm:"comment:返回首页地址" json:"redirect"` //返回首页地址
	// Children  []Menus //
	Component string `gorm:"comment:组件" json:"component"` //组件
	Icon      string `gorm:"comment:图标" json:"icon"`      //图标
	Title     string `gorm:"comment:标题" json:"title"`     //图标
	Name      string `gorm:"comment:页面名称" json:"name"`    //页面名称
	// Meta      Meta   //
	ParentId int64 `gorm:"comment:父节点id" json:"parentId"` //父节点id
}
