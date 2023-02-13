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
		metrics := analyze(pkg)
		fmt.Println(Metrics_toString(metrics))
	}
}
