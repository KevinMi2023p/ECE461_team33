package package_analyzer

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	pkgs := []string{
		"xml2js",
		"bluebird",
		"aws-sdk",
		"ts-node",
		"tar",
		"fake",
	}

	for _, pkg := range pkgs {
		metric := analyze(pkg)
		
		fmt.Println(metric)
	}
}
