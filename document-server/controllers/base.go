package controllers

import (
	"net/http"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

// args[0] http status, args[1] http error message
func (base *BaseController) ErrorHandler(err error, args ...interface{}) {
	if err == nil {
		return
	}
	status := http.StatusBadRequest
	message := err.Error()

	switch len(args) {
	case 1:
		status = args[0].(int)
	case 2:
		status = args[0].(int)
		message = args[1].(string)
	}

	base.CustomAbort(status, message)
}
