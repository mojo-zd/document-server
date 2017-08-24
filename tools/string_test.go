package tools

import (
	"fmt"
	"testing"
)

func Test_Random(t *testing.T) {
	fmt.Println(RandomString(RANGE_32))
	fmt.Println(RemoveSuffix("static/upload/mmm.xls", []Remove{{Split: "/"}, {Split: ".", Reverse: true}}))
}
