package user

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mahdi/Newsgram/common"
	"github.com/mahdi/Newsgram/models"
	"github.com/mahdi/Newsgram/models/dto"
	"github.com/mahdi/Newsgram/role"
	"github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"
)

type AuthorService interface {
	SignUp(ctx *gin.Context)
	SignIn(ctx *gin.Context)
	SignOut(ctx *gin.Context)
	VlidateToken(uid string, token string) bool
	//ProtectedEndpoint(ctx *gin.Context)
	FindByUid(uid string) *models.User
	//Fetch(ctx *gin.Context)
}
type authorService struct {
	Repo     UserRepository
	RoleRepo role.RoleRepository
}

func NewUserService(Repo UserRepository, RoleRepo role.RoleRepository) AuthorService {
	return &authorService{Repo: Repo, RoleRepo: RoleRepo}
}

func (this *authorService) SignUp(ctx *gin.Context) {
	var signupForm dto.SignUpForm
	_ = ctx.Bind(&signupForm)
	//logrus.Println(signupForm)
	//ctx.ReadJSON(&signupForm)
	//logrus.Println(signupForm)

	bool, _ := govalidator.ValidateStruct(&signupForm)

	if bool {
		var count = this.Repo.ExistByUsernameAndPassword(signupForm.Username, signupForm.Password)
		if count == false {
			ctx.JSON(200, dto.BaseDTO{
				Status:   200,
				Message:  common.KEY_USERNAME_PASSWORD_DUPLICATE,
				ObjectId: "",
				Active:   false,
			})
			return
		}

		var fullUser models.User
		fullUser.Username = signupForm.Username
		fullUser.Password = signupForm.Password
		fullUser.DisplayName = signupForm.DisplayName
		fullUser.Active = signupForm.Active

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": signupForm.Username,
			"password": signupForm.Password,
		})
		tokenString, error := token.SignedString([]byte("secret"))
		if error != nil {
			fmt.Println(error)
		}
		fullUser.Token = tokenString

		fullUser.FirstName = signupForm.FirstName
		fullUser.LastName = signupForm.LastName
		fullUser.Adderss = signupForm.Adderss
		var role models.Role
		var roles = []*models.Role{}
		if signupForm.TypeUser == 1 { // Admin

			role, _ = this.RoleRepo.First("Admin")
			roles[0] = &models.Role{Id: role.Id, Name: role.Name, User: role.User}
			fullUser.Roles = roles
			this.Repo.Create(fullUser)

		} else if signupForm.TypeUser == 2 { //User

			role, _ = this.RoleRepo.First("User")
			roles[0] = &models.Role{Id: role.Id, Name: role.Name, User: role.User}
			fullUser.Roles = roles
			this.Repo.Create(fullUser)

		} else if signupForm.TypeUser == 3 { // customer

			role, _ = this.RoleRepo.First("Customer")
			roles[0] = &models.Role{Id: role.Id, Name: role.Name, User: role.User}
			fullUser.Roles = roles
			this.Repo.Create(fullUser)
		}
		ctx.JSON(200, dto.BaseDTO{
			Status:  200,
			Message: common.KEY_SIGNUP_SUCCESSFULLY,
			Active:  true,
		})
	} else {
		ctx.JSON(200, dto.BaseDTO{
			Status:   200,
			Message:  common.KEY_NOT_VALID_JSON,
			ObjectId: "",
			Active:   false,
		})
	}

}
func (this *authorService) FindByUid(uid string) *models.User {
	account := this.FindByUid(uid)
	if account != nil {
		return account
	}
	return nil
}
func (this *authorService) SignIn(ctx *gin.Context) {
	var account dto.LoginForm
	ctx.Bind(&account)
	fullAccount, e := this.Repo.FindByUsernameAndPassword(account.Username, account.Password)
	if e == nil {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": account.Username,
			"password": account.Password,
		})
		tokenString, error := token.SignedString([]byte("secret"))
		if error != nil {
			fmt.Println(error)
		}

		logrus.Infoln(fullAccount)

		this.Repo.UpdateTokenAccount(fullAccount.Id, tokenString)
		ctx.JSON(200, dto.BaseDTO{
			Status:  200,
			Message: common.KEY_SIGNIN_SUCCESSFULLY,
			Active:  true,
			Token:   tokenString,
		})
	} else {
		ctx.JSON(200, dto.BaseDTO{
			Status:   200,
			Message:  common.KEY_USERNAME_PASSWORD_INCORRECT,
			ObjectId: "",
			Active:   false,
			Token:    "",
		})
	}

}
func (this *authorService) SignOut(ctx *gin.Context) {

	var jwtForm dto.BaseDTO
	ctx.Bind(&jwtForm)

	token, _ := jwt.Parse(jwtForm.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return []byte("secret"), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var account models.User
		mapstructure.Decode(claims, &account)
		if account.Username != "" {
			acc, e := this.Repo.FindById(account.Id)
			if e != nil {
				ctx.JSON(200, dto.BaseDTO{
					Status:   200,
					Message:  common.KEY_USERNAME_PASSWORD_INCORRECT,
					ObjectId: "",
					Active:   false,
					Token:    "",
				})
			} else {
				this.Repo.SingOutUser(acc.Id)
				ctx.JSON(200, dto.BaseDTO{
					Status:   200,
					Message:  common.KEY_SIGNOUT,
					ObjectId: "",
					Active:   false,
					Token:    "",
				})
			}
		}
	} else {
		ctx.JSON(200, dto.BaseDTO{
			Status:   200,
			Message:  common.KEY_INVALID_AUTHORIZATION_TOKEN,
			ObjectId: "",
			Active:   false,
		})
	}

}

func (this *authorService) VlidateToken(uid string, token string) bool {

	fullaccount, _ := this.Repo.FindByIdAndToken(uid, token)
	if fullaccount.Username != "" {
		token, _ := jwt.Parse(fullaccount.Token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error")
			}
			return []byte("secret"), nil
		})
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			var account models.User
			mapstructure.Decode(claims, &account)
			if account.Id == fullaccount.Id && account.Username == fullaccount.Username && account.Password == fullaccount.Password {
				return true

			} else {
				return false
			}
		} else {
			return false
		}

	} else {
		return false
	}
}
func (this *authorService) ProtectedEndpoint(ctx *gin.Context) {
	params, _ := ctx.Params.Get("token") // request.URL.Query()
	//ctx.ReadJSON(&jwtForm)

	token, _ := jwt.Parse(params, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return []byte("secret"), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var user models.User
		mapstructure.Decode(claims, &user)
		ctx.JSON(200, user)
	} else {
		ctx.JSON(200, dto.BaseDTO{
			Status:   200,
			Message:  "Invalid authorization token",
			ObjectId: "",
			Active:   false,
		})
	}
}

func (this *authorService) Fetch(ctx *gin.Context) {
	//accounts, _ := this.Repo.Fetch()
	//ctx.JSON(200,accounts)
}
