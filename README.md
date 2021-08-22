# PIN code generator (Go Backend)

## Prerequisites
Will need to have GO installed
- [GO](https://golang.org/doc/tutorial/getting-started)

## Get Started
```
go get -u github.com/abdulpapalek/pin
# OR
copy the pin/ dir in local project 
```
### See sample below for usage.
```
package main

import (
	"fmt"

	"github.com/abdulpapalek/pin"
)

func main() {
	arr := pin.GenerateBatchPIN()

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
```
