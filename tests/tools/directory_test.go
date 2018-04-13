package tools

import (
	"testing"

	"github.com/astaxie/beego/utils"
	"github.com/document-server/tool"
)

func Test_Create(t *testing.T) {
	err := tool.NewDirectory().Create("mojo/test")
	utils.Display("错误信息", err)
	utils.Display("", tool.NewDirectory().Delete("mojo"))
}
