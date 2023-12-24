// Copyright 2023 LegendCastle. All rights reserved.
// MIT License
package lc

import (
	"fmt"
	"reflect"
	"strings"
)

type LCStr struct {
	fmt.Stringer
	_str           string
	_nextArgNumber int
}

func NewStr(fmtstr string) *LCStr {
	return &LCStr{
		_str:           fmtstr,
		_nextArgNumber: 0,
	}
}

func _guessFmtChar(variable interface{}, extraFmt ...interface{}) string {
	t := reflect.TypeOf(variable)
	k := t.Kind()
	f := ""

	switch k {
	case reflect.String:
		f = "%s"
	case reflect.Bool:
		f = "%t"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		f = "%d"
	case reflect.Float32, reflect.Float64:
		if len(extraFmt) > 0 {
			f = "%." + fmt.Sprintf("%d", extraFmt[0]) + "f"
		} else {
			f = "%f"
		}
	case reflect.Map:
		f = "%v"
	default:
		f = "%+v"

	}
	return f
}

func (lcstr *LCStr) String() string {
	return lcstr._str
}

func (lcstr *LCStr) Arg(variable interface{}, extraFmt ...interface{}) *LCStr {
	argLoc := strings.Index(lcstr._str, "%")
	if argLoc != -1 && variable != nil {
		tobeRep := fmt.Sprintf("%%%d", lcstr._nextArgNumber+1)
		replacFmt := _guessFmtChar(variable, extraFmt...)
		replac := fmt.Sprintf(replacFmt, variable)
		lcstr._str = strings.Replace(lcstr._str, tobeRep, replac, 1)
		lcstr._nextArgNumber = lcstr._nextArgNumber + 1
	}
	return lcstr
}
