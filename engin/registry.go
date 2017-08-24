package engin

import (
	"github.com/astaxie/beego/orm"
	"github.com/mojo-zd/document-server/models"
)

func init() {
	orm.RegisterModelWithPrefix("t_", new(models.User))
}
