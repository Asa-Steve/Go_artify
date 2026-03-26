package main

import (
	"fmt"
	"os"

	"github.com/Asa-Steve/Go_artify/artify"
)

func main() {

	// ensuring exactly one argument is passed and nothing else
	if len(os.Args) != 2 {
		fmt.Println("Error: invalid number of argument")
		return
	}

	// initiating the program
	err := artify.Init(os.Args[1], os.Stdout)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

}
