package common

import (
	"github.com/astaxie/beego/orm"
	"github.com/mahdi/News/models"
	"github.com/sirupsen/logrus"
)

func AddSubject(db orm.Ormer) {
	table := db.QueryTable(models.Subject{})
	count, _ := table.Count()
	logrus.Println(count)
	if count <= 0 {

		subject := make([]models.Subject, 12)
		subject[0] = models.Subject{Title: "سیاست"}
		subject[1] = models.Subject{Title: "اقتصاد"}
		subject[2] = models.Subject{Title: "فرهنگ"}
		subject[3] = models.Subject{Title: "ورزش"}
		subject[4] = models.Subject{Title: "حوادث"}
		subject[5] = models.Subject{Title: "اجتماع"}
		subject[6] = models.Subject{Title: "فناوری"}
		subject[7] = models.Subject{Title: "سلامت"}
		subject[8] = models.Subject{Title: "جهان"}
		subject[9] = models.Subject{Title: "عکس"}
		subject[10] = models.Subject{Title: "ویدئو"}
		subject[11] = models.Subject{Title: "کتاب"}

		multi, _ := db.InsertMulti(len(subject), subject)
		logrus.Println(multi)
	}

}

func AddUserAdmin(db orm.Ormer) {

	var username = "mahdihp"
	var password = "110110"
	var user models.User
	user_tbl := db.QueryTable("users")

	_ = user_tbl.Filter("username", username).Filter("password", password).One(&user)

	logrus.Println(user.Username)
	if user.Username == "" && user.Password == "" {
		var newUser = models.User{Username: "mahdi", Password: "110110"}
		index, _ := db.Insert(&newUser)
		logrus.Println("Inser Admin User... ", index)
	}

	/*var count = -1
	db.Table(models.User.TableName(models.User{})).Count(&count)

	//log.Infoln("Record Account Count: ", count)

	if count <= 0 {
		if db.HasTable(models.User{}) == false {
			db.Table(models.User.TableName(models.User{})).CreateTable(&models.User{})
		}


		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": username,
			"password": password,
		})
		tokenString, error := token.SignedString([]byte("secret"))
		if error != nil {
			fmt.Println(error)
		}

		var account = &models.User{Username: &username, Password: &password, Token: tokenString, Active: true,
			Roles: []models.Role{role}}

		log.Println("Insert Record User Rows Affected :", db.Create(&account).RowsAffected)
	}*/
}

func AddRole(db orm.Ormer) {
	roles := make([]models.Role, 3)
	table := db.QueryTable(models.Role{})

	roles[0] = models.Role{Name: "Admin"}
	roles[1] = models.Role{Name: "Author"}
	roles[2] = models.Role{Name: "User"}
	count, _ := table.Count()
	if count != 3 {
		multi, _ := db.InsertMulti(len(roles), roles)
		logrus.Println(multi)
	}
}
