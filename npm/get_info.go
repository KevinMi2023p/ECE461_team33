package npm

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// used to make the request string
const npm_registry_url_part string = "https://registry.npmjs.org/%s"

// alias of map[string]any (same as map[string]interface{}) because typing that is annoying
type NpmInfo = map[string]any;

// performs the get request and parses the json
func Get_NpmInfo(pkg string) *NpmInfo {
	// get request url
	requestUrl := fmt.Sprintf(npm_registry_url_part, pkg)

	// send get request
	response, responseError := http.Get(requestUrl)

	if (responseError != nil) {
		return nil
	}

	// read response body
	defer response.Body.Close()
	bodyBytes, readError := io.ReadAll(response.Body)

	if (readError != nil) {
		return nil
	}

	var data *NpmInfo = new(NpmInfo)

	// parse json from the response body
	jsonError := json.Unmarshal(bodyBytes, data)

	if (jsonError != nil) {
		return nil
	}

	return data
}

// returns the value mapped to key, if info is *map[string]any. otherwise, returns nil
func Get_value_from_info(info any, key string) any {
	if (info == nil) {
		return nil
	}

	switch info.(type) {
		
		case NpmInfo:
			value, noError := info.(NpmInfo)[key]
			if (noError) {
				return value
				} else {
					return nil
				}
				
		default:
			return nil
	
	}
}

// returns info's value at keys, since info is essentially a nested map with string keys
func Get_nested_value_from_info(info *NpmInfo, keys []string) any {
	if (info == nil) {
		return nil
	}

	var result any = *info

	for i := 0; i < len(keys); i++ {
		result = Get_value_from_info(result, keys[i])
	}

	return result
}
