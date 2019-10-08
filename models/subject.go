package models

type Subject struct {
	Id      int      `orm:"auto"`
	Title   string   `orm:"type:varchar(255)"`
	ENTitle string   `orm:"type:varchar(255)"`
	Details string   `orm:"type:varchar(255)"`
	Parent  *Subject `orm:"rel(one);null"`
	News    *News    `orm:"reverse(one)"`
}

func (u *Subject) TableName() string {
	return "subjects"
}
