package models

import "time"

type Comment struct {
	Id      uint      `orm:"auto"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
	Text    string
	News    *News `orm:"rel(fk)"`
}

func (u *Comment) TableName() string {
	return "comments"
}
