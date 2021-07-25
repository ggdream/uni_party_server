package main

import (
	"fmt"
	"os"
)

func main() {
	if err := New(os.Args[1]).Create(); err != nil {
		fmt.Println(err.Error())
		return
	}
}
