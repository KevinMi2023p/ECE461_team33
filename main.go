package main

import (
	"fmt"
	"os"

	"github.com/KevinMi2023p/ECE461_TEAM33/urlprogramfiles" //https://linguinecode.com/post/how-to-import-local-files-packages-in-golang
)

// main function will handle the command line arguments
func main() {
	if len(os.Args) != 2 {
		fmt.Println("Incorrect number of arguments")
		os.Exit(1)
	}

	var argument string = os.Args[1]
	if argument == "install" {

	} else if argument == "build" {

	} else if argument == "test" {

	} else if urlprogramfiles.Check_valid_url(argument) {

	} else {
		fmt.Println("Invalid arguments given")
		os.Exit(1)
	}

}
