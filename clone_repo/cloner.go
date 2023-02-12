package main

//git url parser
import (
	"fmt"
	"os/exec"

	giturl "github.com/KevinMi2023p/ECE461_team33"
)

func is_git_url(repo string) bool {
	_, errror := giturl.NewGitURL(repo) // parse URL, returns error if none git url

	return errror == nil
}

// CloneRepo clones a github repo
func Clone_repo(args []string) {

	//github repo URL
	repo := args[0]

	//verify that is an actual github repo URL
	if !is_git_url(repo) {
		// return
		fmt.Print("Not Valid Git Url ")
	}
	//clone directory
	cmd.Dir := "/Users/bigsteve/ECE461_team33-1/clone_repo/tester"
	//Clones Github Repo
	cmd := exec.Command("git", "clone", repo)
	err := cmd.Run()
	if err != nil {
		fmt.Print("Something went wrong ")
		// something went wrong
	}

}