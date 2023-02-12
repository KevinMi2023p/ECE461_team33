package bus_factor

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"github.com/KevinMi2023p/ECE461_TEAM33/npm"
)

// alias the type because I'm lazy
type NpmInfo = npm.NpmInfo

// uses the cloned repository to determine the bus size 
func Get_minimum_bus_size(gitPath string) int {
	// analyze the cloned repository at gitPath
	cmd := exec.Command("python", "bus_factor.py", fmt.Sprintf("\"%s\"", gitPath))
	output, err := cmd.Output()

	if (err != nil) {
		return 0
	}

	// parse bus_size from python output
	i, parseError := strconv.Atoi(strings.TrimSpace(string(output)))

	if (parseError != nil) {
		return 0
	}

	return i
}

// calculates a bus factor (between 0 and 1) from the bus size
func calculate_bus_factor(bus_size int) float32 {
	if (bus_size < 1) {
		return 0.0
	}

	return (float32(bus_size) - 1) / float32(bus_size)
}

// to be called externally, once the repo has been cloned. Don't put white space in gitPath, it will break
func Get_bus_factor(gitPath string) float32 {
	bus_size := Get_minimum_bus_size(gitPath)
	return calculate_bus_factor(bus_size)
}
