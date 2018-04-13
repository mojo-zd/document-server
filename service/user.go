package service

import (
	"github.com/astaxie/beego/orm"
	"github.com/document-server/models"
)

type UserService struct {
	BaseService
}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) Query() (users []*models.User, err error) {
	users = []*models.User{}
	_, err = orm.NewOrm().QueryTable(new(models.User)).Filter("Deleted", false).All(&users)
	return
}
