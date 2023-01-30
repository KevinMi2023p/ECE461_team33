package npm_registry

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const npm_registry_url_part string = "https://registry.npmjs.org/%s"

func npm_registry_get_json(pkg string) map[string]interface{} {
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
