package main

import (
	"fmt"
	"rcscan"
)

func main() {

	rc, err := rcscan.New("./example.rc")
	if err != nil {
		fmt.Println(err)
		return
	}

	// get param A from empty section:

	param, err := rc.Get("", "paramA")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("paramA:", param)

	// get param A from specified section:

	param, err = rc.Get("Section 1", "paramA")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("paramA:", param)

}
