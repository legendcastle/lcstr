Licence:
-----------
MIT

Download:
-----------
go get -u github.com/legendcastle/lcstr

Usage:
-----------

```
package main

import (
	"fmt"
	"time"

	"github.com/legendcastle/lcstr/lc"
)

func main() {
	str := lc.NewStr("My name is %1. My age is %2. My Height is %3 meter. Today is %4").Arg("John Smith").Arg(18).Arg(1.78, 2).Arg(time.Now().Format("2006-01-02"))
	fmt.Println(str)
}
```