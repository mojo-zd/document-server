package engine

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	"github.com/document-server/models"
)

var (
	name    = "default" // 数据库别名
	force   = false     // drop table 后再建表
	verbose = false     // 打印执行过程
)

func init() {
	orm.RegisterModel(new(models.User))

	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
}
