package rampuptime

import (
	"fmt"
	"testing"

	"github.com/KevinMi2023p/ECE461_TEAM33/npm"
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
		info := npm.Get_NpmInfo(pkg)
		fmt.Println("Package used:\t" + pkg)

		if info == nil {
			fmt.Println("Couldn't get npm registry")
		} else {
			fmt.Println(Ramp_up_score(info))
			// fmt.Println((*info)["readme"])
		}
	}
}
