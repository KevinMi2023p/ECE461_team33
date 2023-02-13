package responsiveness

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	repos := []string{ "https://github.com/KevinMi2023p/ECE461_TEAM33/" }
	token := os.Getenv("GITHUB_TOKEN")

	fmt.Printf("GitHub token:\t%s\n\n", token)
	
	for _, repo := range repos {
		issues := Get_issues(repo, token)
		fmt.Printf("Responsiveness for %s:\t%f\n", repo, Responsiveness(issues))
	}
}
