package tools

import (
	"fmt"
	"testing"

	"github.com/document-server/tool"
)

func Test_UUID(t *testing.T) {
	fmt.Println(tool.NewUUID())
}

func Test_Suffix(t *testing.T) {
	fmt.Printf(tool.Suffix("name.excel", "."))
}

func Test_Substring(t *testing.T) {
	fmt.Printf(tool.Substring("name.excel", "."))
}
