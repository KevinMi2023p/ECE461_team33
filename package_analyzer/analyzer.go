package package_analyzer

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"
	
	. "github.com/KevinMi2023p/ECE461_TEAM33/bus_factor"
	. "github.com/KevinMi2023p/ECE461_TEAM33/npm"
	. "github.com/KevinMi2023p/ECE461_TEAM33/ramp_up_time"
	. "github.com/KevinMi2023p/ECE461_TEAM33/responsiveness"
	. "github.com/KevinMi2023p/ECE461_TEAM33/license_compatibility"
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

var Npm_package_regex *regexp.Regexp = regexp.MustCompile(`^https://www\.npmjs\.com/package/([\w-\.]+)$`)
var Github_package_regex *regexp.Regexp = regexp.MustCompile(`^https://github.com/([\w-]+)/([\w-]+)$`)

var Package_github_url_regex *regexp.Regexp = regexp.MustCompile(`git\+(https://|ssh://git@)(github\.com/[\w-]+/[\w-]+)\.git$`)

const Npm_registry_url_part string = "https://registry.npmjs.org/%s"
const Github_url_part string = "https://github.com/%s/%s"

const Github_api_part string = "https://api.github.com/repos/%s/%s"

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
			if (!Package_github_url_regex.MatchString(repoUrl)) {
				return nil
			}
			
			submatches := Package_github_url_regex.FindStringSubmatch(repoUrl)
			if (submatches == nil || len(submatches) < 3) {
				return nil
			}

			// return the result as a pointer
			var result *string = new(string)
			*result = "https://" + submatches[2]

			return result
		}
	}

	return nil
}

// calculates weighted net score from other metrics
func net_score(metrics *Metrics) float32 {
	return (metrics.net_score + metrics.bus_factor + metrics.ramp_up_time + metrics.responsiveness + metrics.correctness) / 5
}

func analyze(url string) *Metrics {
	var metrics *Metrics
	var submatches []string
	var token string
	var repo_api string
	var issues *[]RepoIssue = nil
	var client *http.Client

	if (Github_package_regex.MatchString(url)) {
		// get the safe url
		submatches = Github_package_regex.FindStringSubmatch(url)
		if (submatches == nil || len(submatches) < 3) {
			return nil
		}

		metrics = new(Metrics)
		metrics.url = fmt.Sprintf(Github_url_part, submatches[1], submatches[2])
	
		// client
		client = &http.Client{}

		// github repo api
		repo_api = fmt.Sprintf(Github_api_part, submatches[1], submatches[2])

		// repo issues
		token = os.Getenv("GITHUB_TOKEN")
		issues = Get_issues(repo_api, token, client)

		// responsiveness
		metrics.responsiveness = Responsiveness(issues)

		// bus factor
		metrics.bus_factor = Get_bus_factor(metrics.url)

		// ramp up time
		metrics.ramp_up_time = Ramp_up_score_github(repo_api, token, client)
	
		// correctness
		metrics.correctness = 0
	
		// net score
		metrics.net_score = net_score(metrics)
	} else if (Npm_package_regex.MatchString(url)) {
		// get the safe url
		submatches = Npm_package_regex.FindStringSubmatch(url)
		if (submatches == nil || len(submatches) < 2) {
			return nil
		}

		metrics = new(Metrics)
		metrics.url = url

		// bus factor
		metrics.bus_factor = 0
		
		// get npm info
		reg_url := fmt.Sprintf(Npm_registry_url_part, submatches[1])
		info := Get_NpmInfo(reg_url)
		
		// get the github url
		githubUrl := github_url(info)

		if (githubUrl != nil && Github_package_regex.MatchString(*githubUrl)) {
			// client
			client = &http.Client{}

			// get the safe url
			submatches = Github_package_regex.FindStringSubmatch(*githubUrl)
			
			// github repo api
			repo_api = fmt.Sprintf(Github_api_part, submatches[1], submatches[2])
			
			// repo issues
			token = os.Getenv("GITHUB_TOKEN")
			issues = Get_issues(repo_api, token, client)

			// bus factor
			metrics.bus_factor = Get_bus_factor(*githubUrl)
		}
		
		// responsiveness
		metrics.responsiveness = Responsiveness(issues)
	
		// ramp up time
		metrics.ramp_up_time = Ramp_up_score_npm(info)
	
		// correctness
		metrics.correctness = 0

		// license
		metrics.license = License_compatibity(info)
	
		// net score
		metrics.net_score = net_score(metrics)
	} else {
		metrics = nil
	}

	return metrics
}
