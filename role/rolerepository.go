package role

import (
	"github.com/astaxie/beego/orm"
	"github.com/mahdi/Newsgram/models"
)

type RoleRepository interface {
	Fetch() ([]*models.Role, bool)
	First(Name string) (models.Role, bool)
}

func NewRoleRepository(instance orm.Ormer) RoleRepository {
	return &roleRepository{Instance: instance}
}

type roleRepository struct {
	Instance orm.Ormer
}

func (this *roleRepository) First(Name string) (models.Role, bool) {
	var role models.Role
	role_tbl := this.Instance.QueryTable("roles")
	_ = role_tbl.Filter("Name", Name).One(&role)
	return role, true
}

func (this *roleRepository) Fetch() ([]*models.Role, bool) {
	roles := make([]*models.Role, 0)
	role_tbl := this.Instance.QueryTable("roles")
	_, _ = role_tbl.All(&roles)
	//if len(roles) == 0 && err != nil {
	//	logrus.Error(err)
	//	return nil, nil
	//}
	return roles, true
}
