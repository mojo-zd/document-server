package engin

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

var (
	name    = "default" // 数据库别名
	force   = true      // drop table 后再建表
	verbose = true      // 打印执行过程
)

func init() {
	orm.RegisterModel()

	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
}
