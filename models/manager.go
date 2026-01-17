package models

type Manager struct {
	Id int
	Username string
	Password string
	Mobile string
	Email string
	Status int
	RoleId int
	AddTime int
	IsSuper int
	Role Role  `gorm:"foreignKey:RoleId;references:Id"` // 关联字段
}

func (Manager) TableName() string {
	return "manager"
}