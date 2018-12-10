package main

import (
	"fmt"
	"os"
)

//error handeling
func check(err error, ext int) {
	if err != nil && ext == 1 {
		fmt.Println(err)
		os.Exit(1)
	} else if err != nil && ext == 0 {
		fmt.Println(err)
	}
}
