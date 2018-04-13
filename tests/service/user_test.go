package service

import (
	"testing"

	"github.com/astaxie/beego/utils"
	"github.com/document-server/models"
	"github.com/document-server/service"
)

var userService = service.NewUserService()

func Test_Get_User(t *testing.T) {
	user := &models.User{Name: "mojo"}
	err := userService.Read(user, "Name")
	utils.Display("=== error info is ===", err)
	utils.Display("=== user is ===", user)
}

func Test_Insert_User(t *testing.T) {
	user := &models.User{Name: "mojo"}
	err := userService.InsertOrUpdate(user)
	utils.Display("===error info is ===", err)
}

func Test_Insert_Multi_User(t *testing.T) {
	users := []*models.User{{Name: "mt"}}
	err := userService.InsertMulti(&users)
	utils.Display("error info is ", err)
	utils.Display("result is ", users)
}

func Test_Query_User(t *testing.T) {
	users, err := userService.Query()
	utils.Display("error info is ", err)
	utils.Display("result is ", users)
}

func Test_Update_User(t *testing.T) {
	user := &models.User{Id: 1, Name: "mojo", Password: "xxxx"}
	err := userService.Update(user, "Name")
	utils.Display("error info is ", err)
}

func Test_Update_User_With_Map(t *testing.T) {
	user := map[string]interface{}{"Name": "mojo"}
	err := userService.UpdateWithMap(user, models.User{})
	utils.Display("error info is ", err)
}
