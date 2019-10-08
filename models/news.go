package models

import "time"

type News struct {
	Id        uint    `orm:"auto"`
	Newscode  *string `orm:"unique;null"`
	Title     string
	Summary   string
	Text      string
	Tag       string
	Reference string
	MediaLink string    `orm:"type(json)"`
	Source    string    `orm:"type(json)"`
	Subject   *Subject  `orm:"null;rel(one);on_delete(set_null)"`
	Created   time.Time `orm:"auto_now_add;type(datetime)"`
	Updated   time.Time `orm:"auto_now;type(datetime)"`
	User      *User     `orm:"rel(fk)"`
	//Comment   []Comment `gorm:"foreignkey:NewsId"`
}

func (u *News) TableName() string {
	return "news"
}
