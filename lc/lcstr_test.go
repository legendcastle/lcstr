package lc

import (
	"fmt"
	"testing"
)

func TestLCStr(t *testing.T) {
	str := NewStr("My name is %1. My age is %2. My Height is %3 meter.").Arg("John Smith").Arg(18).Arg(1.78, 2)
	fmt.Println(str)
	t.Log(str)
}
