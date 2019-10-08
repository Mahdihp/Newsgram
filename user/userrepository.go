package user

import (
	"github.com/astaxie/beego/orm"
	"github.com/mahdi/Newsgram/models"
	"github.com/sirupsen/logrus"
	"sync"
)

type UserRepository interface {
	Create(account models.User)
	Fetch() (accounts []models.User, err error)
	FindByUsernameAndPassword(username string, password string) (models.User, error)
	ExistByUsernameAndPassword(username string, password string) bool
	FindByIdAndToken(id string, token string) (models.User, error)
	//CountByUserAndPassAndNati(username string, password string, nationalcode string) int
	FindByIdAndUsernameAndPassword(id string, username string, password string) (models.User, error)
	UpdateTokenAccount(id uint, token string)
	SingOutUser(id uint)
	FindById(id uint) (models.User, error)
	//MaxAccountNumber() (int64, error)
}

func NewUserRepository(instance orm.Ormer) UserRepository {
	return &userRepository{Instance: instance}
}

type userRepository struct {
	Instance orm.Ormer
	mu       sync.RWMutex
}

func (this *userRepository) UpdateTokenAccount(id uint, token string) {
	var user = models.User{}
	user_tbl := this.Instance.QueryTable("users")
	e := user_tbl.Filter("id", id).One(&user)
	//err := this.Instance.Where("id = ?", id).First(&user).Error
	if e == nil {
		user.Token = token
		this.Instance.Update(&user)
	}
}

func (this *userRepository) FindByIdAndUsernameAndPassword(id string, username string, password string) (models.User, error) {
	var user = models.User{}
	user_tbl := this.Instance.QueryTable("users")
	user_tbl.Filter("id", id).Filter("username", username).Filter("password", password).One(&user)
	return user, nil
}

func (this *userRepository) Fetch() ([]models.User, error) {
	user := make([]models.User, 0)
	//err := this.Instance.Find(&user).Error
	user_tbl := this.Instance.QueryTable("users")
	_, err := user_tbl.All(&user)
	//err := this.Instance.Preload("Role").Find(&user).Error
	if len(user) == 0 {
		logrus.Error(err)
		return nil, nil
	}
	return user, nil
}

func (this *userRepository) SingOutUser(id uint) {
	var user = models.User{}
	user_tbl := this.Instance.QueryTable("users")
	err := user_tbl.Filter("id", id).One(&user)
	//err := this.Instance.Where("id = ?", id).First(&user).Error
	if err == nil {
		user.Token = ""
		this.Instance.Update(&user)
	}
}
func (this *userRepository) FindById(id uint) (models.User, error) {
	var user = models.User{}
	user_tbl := this.Instance.QueryTable("users")
	user_tbl.Filter("id", id).One(&user)
	//this.Instance.Where("id = ?", id).First(&user)
	//if err != nil {
	//	logrus.Error(err)
	//	return nil, nil
	//}
	return user, nil
}

func (this *userRepository) FindByIdAndToken(id string, token string) (models.User, error) {
	var user = models.User{}
	user_tbl := this.Instance.QueryTable("users")
	user_tbl.Filter("id", id).Filter("token", token).One(&user_tbl)
	//this.Instance.Where("id = ? AND token = ? ", id, token).First(&user)
	//if err != nil {
	//	logrus.Error(err)
	//	return nil, nil
	//}
	return user, nil
}
func (this *userRepository) ExistByUsernameAndPassword(username string, password string) bool {
	var user = models.User{}
	user_tbl := this.Instance.QueryTable("users")
	user_tbl.Filter("username", username).Filter("password", password).One(&user)
	//this.Instance.Where("username = ? AND password = ? ", username, password).First(&user)
	if user.Username != "" {
		//	logrus.Error(err)
		return true
	} else {
		return false
	}
}

func (this *userRepository) FindByUsernameAndPassword(username string, password string) (models.User, error) {
	var user = models.User{}
	user_tbl := this.Instance.QueryTable("users")
	user_tbl.Filter("username", username).Filter("password", password).One(&user) //if err != nil {
	//	logrus.Error(err)
	//	return nil, nil
	//}
	return user, nil
}

func (this *userRepository) Create(user models.User) {

}
