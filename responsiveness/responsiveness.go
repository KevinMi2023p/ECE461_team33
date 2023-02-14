package responsiveness

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"github.com/KevinMi2023p/ECE461_TEAM33/npm"
)

// used to make the request string
const github_issues_url_part string = "%s/issues?filter=all&state=all"
const bearer_auth_part string = "Bearer %s"

// alias of map[string]any (same as map[string]interface{}) because typing that is annoying
type RepoIssue = map[string]any;

// performs the get request and parses the json
func Get_issues(repo_api string, token string) *[]RepoIssue {
	requestUrl := fmt.Sprintf(github_issues_url_part, repo_api)
	auth := fmt.Sprintf(bearer_auth_part, token)

	// create new request
	request, requestError := http.NewRequest("GET", requestUrl, nil)

	if (requestError != nil) {
		// fmt.Print("Request Error:\t")
		// fmt.Println(requestError)
		return nil
	}

	// add bearer token to the header
	request.Header.Add("Accept", "application/vnd.github+json")
	request.Header.Add("Authorization", auth)
	request.Header.Add("X-GitHub-Api-Version", "2022-11-28")

	// send request
	client := &http.Client{}
	response, responseError := client.Do(request)

	if (responseError != nil) {
		// fmt.Print("Response Error:\t")
		// fmt.Println(responseError)
		return nil
	}

	// read response body
	defer response.Body.Close()
	bodyBytes, readError := io.ReadAll(response.Body)

	if (readError != nil) {
		// fmt.Print("Read Error:\t")
		// fmt.Println(readError)
		return nil
	}

	if (len(bodyBytes) == 0) {
		return nil
	}

	var data *[]RepoIssue = new([]RepoIssue)

	// parse json from the response body
	jsonError := json.Unmarshal(bodyBytes, data)

	if (jsonError != nil) {
		// fmt.Print("Json Error:\t")
		// fmt.Println(jsonError)
		return nil
	}

	return data
}

// calculate responsiveness from repo issues
func Responsiveness(issues *[]RepoIssue) float32 {
	if (issues == nil) {
		return 0
	}

	bugCount := 0
	closedBugs := 0

	for _, issue := range *issues {
		// check whether the issue is a bug
		labels := npm.Get_value_from_info(issue, "labels").([]interface{})

		for i := 0; i < len(labels); i++ {
			name := npm.Get_value_from_info(labels[i], "name")

			// if this label is "Bug"
			if (name != nil) {
				if (name.(string) == "Bug") {
					i = len(labels)
					bugCount += 1

					// check whether the issue is no longer open
					state := npm.Get_value_from_info(issue, "state")
					if (state != nil) {
						if (state != "open") {
							closedBugs += 1
						}
					}
				}
			}
		}
	}

	if (bugCount > 0) {
		return float32(closedBugs) / float32(bugCount)
	}
	
	return 0.5
}
