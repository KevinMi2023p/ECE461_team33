package npm_registry

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const npm_registry_url_part string = "https://registry.npmjs.org/%s"

type npm_info struct {
	hasRepo bool
	repoUrl string
}

func get_json(pkg string) map[string]interface{} {
	res, gerr := http.Get(fmt.Sprintf(npm_registry_url_part, pkg))

	if (gerr != nil) {
		return nil
	}

	b, rerr := io.ReadAll(res.Body)

	if (rerr != nil) {
		return nil
	}

	var data map[string]interface{}
	jerr := json.Unmarshal(b, &data)

	if (jerr != nil) {
		return nil
	}

	return data
}

func get_value_from_map(i map[string]interface{}, key string) interface{} {
	value, ok := i[key]
	if (ok) {
		return value
	} else {
		return nil
	}
}

func set_repo_from_json(info *npm_info, data map[string]interface{}) {
	info.hasRepo = false
	info.repoUrl = ""

	repoValue := get_value_from_map(data, "repository")

	if (repoValue != nil) {
		repoType := get_value_from_map(repoValue.(map[string]interface{}), "type")

		if (repoType != nil && repoType.(string) == "git") {
			repoString := get_value_from_map(repoValue.(map[string]interface{}), "url")

			if (repoString != nil) {
				url := repoString.(string)

				if (len(url) > 4) {
					info.hasRepo = true
					info.repoUrl = url[4:]
					return
				}
			}
		}
	}
}

func get_info(pkg string) *npm_info {
	data := get_json(pkg)

	if (data == nil) {
		return nil
	}

	var info *npm_info = new(npm_info)

	set_repo_from_json(info, data)

	return info
}
