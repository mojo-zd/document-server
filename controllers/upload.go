package controllers

import (
	"github.com/document-server/constant"
	"github.com/document-server/document"
)

type UploadFile struct {
	BaseController
}
type Uploader struct {
	InitiativeStartRow int `form:"initiativeStartRow"`
	InitiativeColumn   int `form:"initiativeColumn"`
	PassiveStartRow    int `form:"passiveStartRow"`
	PassiveColumn      int `form:"passiveColumn"`
}

// @Title Uploadfile
// @Description create object
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router /file/upload [post]
func (u *UploadFile) UploadFile() {
	loader := &Uploader{}
	u.ParseFromForm(loader)

	passiveFile := u.saveFile("passive")
	initiativeFile := u.saveFile("initiative")
	document.NewDocument(constant.UPLOAD_DIR+passiveFile, constant.UPLOAD_DIR+initiativeFile).SetPassiveColumn(1).SetInitiativeColumn(1).Compare()
}
