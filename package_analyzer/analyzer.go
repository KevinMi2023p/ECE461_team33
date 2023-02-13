package package_analyzer

import (
	"fmt"
	"os"
	"strings"

	. "github.com/KevinMi2023p/ECE461_TEAM33/bus_factor"
	. "github.com/KevinMi2023p/ECE461_TEAM33/npm"
	. "github.com/KevinMi2023p/ECE461_TEAM33/ramp_up_time"
	. "github.com/KevinMi2023p/ECE461_TEAM33/responsiveness"
)

type Metrics struct {
	url string
	bus_factor float32
	ramp_up_time float32
	responsiveness float32
	correctness float32
	license float32
	net_score float32
}

const Metrics_print_format string = "{\"URL\":\"%s\",\"NetScore\":%f,\"RampUp\":%f,\"Correctness\":%f,\"BusFactor\":%f,\"ResponsiveMaintainer\":%f,\"License\":%f}"

func Metrics_toString(metrics Metrics) string {
	return fmt.Sprintf(Metrics_print_format, metrics.url, metrics.net_score, metrics.ramp_up_time, metrics.correctness, metrics.bus_factor, metrics.responsiveness, metrics.license)
}

// get the github url from the json, if there is an associated github
func github_url(info *NpmInfo) *string {
	// get repo type
	repoTypeKeys := []string{ "repository", "type" }
	repoType := Get_nested_value_from_info(info, repoTypeKeys)

	// if the repo is a git repo
	if (repoType != nil && repoType == "git") {
		// get repo address
		repoUrlAny := Get_nested_value_from_info(info, []string{ "repository", "url" })

		if (repoUrlAny != nil) {
			repoUrl := repoUrlAny.(string)

			if (len(repoUrl) > 4) {
				// trim leading "git+", whitespace, and trailing '/' from the url
				repoUrl = strings.Trim(strings.TrimSpace(repoUrl)[4:], "/")

				// remove trailing ".git" from the repo path 
				if (len(repoUrl) > 4) {
					repoUrl = repoUrl[:len(repoUrl) - 4]
				}

				// return the result as a pointer
				var result *string = new(string)
				*result = repoUrl

				return result
			}

		}
	}

	return nil
}

// calculates weighted net score from other metrics
func net_score(metrics Metrics) float32 {
	return metrics.net_score * (metrics.bus_factor + metrics.ramp_up_time + metrics.responsiveness + metrics.correctness) / 4
}

func analyze(pkg string) Metrics {
	metrics := Metrics{}
	metrics.url = fmt.Sprintf(Npm_registry_url_part, pkg)
	
	info := Get_NpmInfo(metrics.url)

	githubUrl := github_url(info)
	if (githubUrl != nil) {
		metrics.url = *githubUrl
	}

	// repo issues
	var issues *[]RepoIssue = nil

	// responsiveness
	metrics.responsiveness = Responsiveness(issues)

	// bus factor
	metrics.bus_factor = 0

	if (githubUrl != nil) {
		token := os.Getenv("GITHUB_TOKEN")
		issues = Get_issues(*githubUrl, token)
		metrics.bus_factor = Get_bus_factor(*githubUrl)
	}

	// ramp up time
	metrics.ramp_up_time = Ramp_up_score(info)

	// net score
	metrics.net_score = net_score(metrics)

	// correctness
	metrics.correctness = 0

	return metrics
}
