package engine

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var dataSource = "%s:%s@/%s?charset=utf8"
var (
	maxIdle = 30
	maxConn = 50
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", fmt.Sprintf(dataSource, "root", "root123", "document"), maxIdle, maxConn)
}
