package models

type Access struct {
  Id int
  ModuleName string // 模块名称
  ActionName string // 操作名称
  Type int // 节点类型: 1.模块  2.菜单 3.操作
  Url string // 路由跳转地址
  ModuleId int // 和当前模型的id关联: module_id=0 表示模块
  Sort int
  Description string
  Status int
  AddTime int
  AccessItem []Access `gorm:"foreignKey:ModuleId;references:Id"` // 关联字段, 当前表和自己是一对多的关系。
  Checked bool `gorm:"-"` // 表示忽略本字段, 表示和数据库表不打交道的。
}

func (Access) TableName() string {
  return "access"
}