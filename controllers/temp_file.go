package controllers

import (
	"github.com/document-server/models"
	"github.com/document-server/tools"
)

type TempFile struct {
	BaseController
}

// @Title Upload File
// @Description create object
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router /file/upload [post]
func (tf *TempFile) UploadFile() {
	file := tf.saveFile("file")
	tempFile := &models.TempFile{UUID: tools.Substring(file, ".")}
	tf.Data["json"] = tempFile
	tf.ServeJSON()
}
