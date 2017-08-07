package controllers

import (
	"fmt"
	"net/http"

	"github.com/document/document-server/constant"
	"github.com/document/document-server/document"
)

type UploadFile struct {
	BaseController
}

// @Title Uploadfile
// @Description create object
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router /file/upload [post]
func (u *UploadFile) UploadFile() {
	passiveFile := u.saveFile("passive")
	initiativeFile := u.saveFile("initiative")
	document.NewDocument(constant.UPLOAD_DIR+passiveFile, constant.UPLOAD_DIR+initiativeFile).SetPassiveColumn(1).SetInitiativeColumn(1).Compare()
}

func (u *UploadFile) saveFile(name string) (fileName string) {
	_, header, err := u.GetFile(name)
	u.ErrorHandler(err)

	fileName = header.Filename
	err = u.SaveToFile(name, constant.UPLOAD_DIR+header.Filename)
	u.ErrorHandler(err, http.StatusUnprocessableEntity, fmt.Sprintf("保存临时文件 %s 失败", header.Filename))

	return
}
