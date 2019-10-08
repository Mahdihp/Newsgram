package models

type Role struct {
	Id   int     `orm:"auto"`
	Name string  `orm:"unique;null;type:varchar(20)"`
	User []*User `orm:"rel(m2m)"`
}

func (Role) TableName() string {
	return "roles"
}
