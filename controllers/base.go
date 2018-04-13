package controllers

import (
	"fmt"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/document-server/constant"
	"github.com/document-server/tools"
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

func (base *BaseController) ParseFromForm(i interface{}) {
	err := base.ParseForm(i)
	base.ErrorHandler(err)
}

func (bc *BaseController) saveFile(name string) (fileName string) {
	_, header, err := bc.GetFile(name)
	bc.ErrorHandler(err)
	fileName = newFileName(header.Filename)
	err = bc.SaveToFile(name, constant.UPLOAD_DIR+fileName)
	bc.ErrorHandler(err, http.StatusUnprocessableEntity, fmt.Sprintf("保存临时文件 %s 失败", header.Filename))

	return
}

func newFileName(name string) string {
	return fmt.Sprintf("%s%s", tools.NewUUID(), tools.Suffix(name, "."))
}
