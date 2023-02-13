package rampuptime

import (
	"regexp"
	"github.com/KevinMi2023p/ECE461_TEAM33/npm"
)

//This module uses the "npm" package that's created locally

func Ramp_up_score(json_data *npm.NpmInfo) float32 {
	//Below is a link of what <json_data> should be
	//https://github.com/npm/registry/blob/master/docs/responses/package-metadata.md

	if json_data == nil {
		// fmt.Println("NPM package hasn't been called yet or doesn't exist")
		return 0
	}

	//Look through map to get README data
	var readme_string = npm.Get_nested_value_from_info(json_data, []string{"readme"})
	if readme_string == nil {
		// fmt.Println("Couldn't read README file")
		return 0
	}

	return calculate_score(readme_string.(string))
}

// Puts string of readme and returns the score
func calculate_score(readme string) float32 {
	reg, err := regexp.Compile("[^A-Za-z0-9]+")
	if err != nil {
		// fmt.Println("Regex failed to compile")
		return 0
	}

	var filter_readme string = reg.ReplaceAllString(readme, "")

	var len_filter int = len(filter_readme)
	var good_amount int = 1000 //if a README has a 1000 character, it should have enough information

	//Condition for perfect score
	if len_filter >= good_amount {
		return 1.0
	}

	return float32(len_filter) / float32(good_amount)
}
