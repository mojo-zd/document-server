package engin

import (
	"github.com/astaxie/beego/orm"
	"github.com/document-server/models"
)

func init() {
	orm.RegisterModelWithPrefix("t_", new(models.User))
}