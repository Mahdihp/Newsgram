package models

import "time"

type User struct {
	Id          uint   `orm:"auto"`
	Username    string `orm:"unique;null;type:varchar(25)"`
	Password    string `orm:"unique;null;type:varchar(25)"`
	FirstName   string `orm:"type:varchar(255)"`
	LastName    string `orm:"type:varchar(255)"`
	DisplayName string `orm:"type:varchar(255)"`
	Email       string `orm:"type:varchar(255)"`
	Website     string
	Education   string
	Adderss     string
	Token       string
	Active      bool
	Created     time.Time `orm:"auto_now_add;type(datetime)"`
	Updated     time.Time `orm:"auto_now;type(datetime)"`

	Roles []*Role `orm:"reverse(many)"`
	News  []*News `orm:"reverse(many)"`
}

func (u User) TableName() string {
	return "users"
}
