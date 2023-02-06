package npm_registry

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// used to make the request string
const npm_registry_url_part string = "https://registry.npmjs.org/%s"

// to be used by other packages; a more useful data structure than a map[string]interface{}
type NpmInfo struct {
	json map[string]interface{}
	repoUrl *string
}

// performs the get request and parses the json to a map/interface
func get_json(pkg string) map[string]interface{} {
	// get request
	res, gerr := http.Get(fmt.Sprintf(npm_registry_url_part, pkg))

	if (gerr != nil) {
		return nil
	}

	// read body
	b, rerr := io.ReadAll(res.Body)

	if (rerr != nil) {
		return nil
	}

	var data map[string]interface{}

	// parse json
	jerr := json.Unmarshal(b, &data)

	if (jerr != nil) {
		return nil
	}

	return data
}

// returns null if the map doesn't contain a value
func Get_value_from_map(i map[string]interface{}, key string) interface{} {
	value, ok := i[key]
	if (ok) {
		return value
	} else {
		return nil
	}
}

// sets repoUrl using the json data
func set_repo_from_json(info *NpmInfo, data map[string]interface{}) {
	info.repoUrl = nil

	repoValue := Get_value_from_map(data, "repository")

	// make sure there's a repo value in the result
	if (repoValue != nil) {
		repoType := Get_value_from_map(repoValue.(map[string]interface{}), "type")

		// check if the repo type is git, that's the only type we're preparing to handle
		if (repoType != nil && repoType.(string) == "git") {
			repoString := Get_value_from_map(repoValue.(map[string]interface{}), "url")

			if (repoString != nil) {
				url := repoString.(string)

				// the url begins with "git+". This should be removed so the url can be used without further manipulation
				if (len(url) > 4) {
					info.repoUrl = new(string)
					*info.repoUrl = url[4:]
					return
				}
			}
		}
	}
}

// perform the get request then read the json into a more useful data structure
func Get_info(pkg string) *NpmInfo {
	data := get_json(pkg)

	var info *NpmInfo = new(NpmInfo)

	info.json = data

	set_repo_from_json(info, data)

	return info
}