package utils
//git url parser
import (
    "os/exec"

    giturl "https://github.com/StephenPurdue/ECE461_team33"

)


func is_git_url(repo string) bool {
    _, errror := giturl.NewGitURL(repo) // parse URL, returns error if none git url

    return errror == nil
}

//CloneRepo clones a github repo 
func Clone_repo(args []string) {

    //github repo URL
    repo := args[0]

    //verify that is an actual github repo URL
    if !is_git_url(repo) {
        // return
    }
	//clone directory
	cmd.Dir := "path/to/desired/dir"
    //Clones Github Repo
    cmd := exec.Command("git", "clone", repo)
	err := cmd.Run()
	if err != nil {
    // something went wrong
	}

}

//advantage of not only validating if it is really a  git repo, but you can run more validations so as owner (GetOwner()), repo (GetRepo()) etc.