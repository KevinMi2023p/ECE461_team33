package main

import (
	"fmt"
	"os"

	//https://linguinecode.com/post/how-to-import-local-files-packages-in-golang
	"github.com/KevinMi2023p/ECE461_TEAM33/installation"
	"github.com/KevinMi2023p/ECE461_TEAM33/maintesting"
	"github.com/KevinMi2023p/ECE461_TEAM33/urlprogramfiles"
)

// main function will handle the command line arguments
func main() {
	if len(os.Args) != 2 {
		fmt.Println("Incorrect number of arguments")
		os.Exit(1)
	}

	var argument string = os.Args[1]
	if argument == "install" {
		installation.Python_pip_install("GitPython")
		installation.Python_pip_install("truckfactor")
	} else if argument == "build" {

	} else if argument == "test" {
		maintesting.MainTest()
	} else if urlprogramfiles.Check_valid_url(argument) {

	} else {
		fmt.Println("Invalid arguments given")
		os.Exit(1)
	}

}
