package controllers

import (
	"net/http"

	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title Login
// @Description Login
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [get]
func (u *UserController) Login() {
	u.Ctx.Request.AddCookie(&http.Cookie{Name: "mojo", Value: "mojo-value"})
	u.ServeJSON()
}
