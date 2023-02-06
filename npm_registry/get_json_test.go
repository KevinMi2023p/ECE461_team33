package npm_registry

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	x := get_json("xml2js")
	for key, _ := range x {
		fmt.Printf("%s\n", key)
	}
	fmt.Println(x["repository"].(map[string]interface{})["type"])
	y := get_info("xml2js")
	fmt.Println(y.repoUrl)
}