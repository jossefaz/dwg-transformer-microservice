package main

import "fmt"

func HandleError(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %s", msg, err)
	}

}
