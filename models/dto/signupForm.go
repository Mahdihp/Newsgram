package dto

type SignUpForm struct {
	DisplayName  string `valid:"required"`
	Username     string `valid:"stringlength(5|20)"`
	Password     string `valid:"stringlength(5|20)"`
	NationalCode string `valid:"stringlength(10|10)"` //0386007551
	MobileNumber string `valid:"stringlength(11|11)"` //09121544275
	Active       bool   `valid:"optional"`
	FirstName    string `valid:"optional"`
	LastName     string `valid:"optional"`
	FatherName   string `valid:"optional"`
	City         string `valid:"optional"`
	Adderss      string `valid:"optional"`
	TypeUser     int    `valid:"required"` //1=Admin & 2=user & 3=customer
}
