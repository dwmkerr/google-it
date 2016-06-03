package main

import (
	"regexp"
	"strconv"
	"strings"
)

// Params represents the parameters which can be provided to the program.
type Params struct {
	ShowHelp struct {
		Present bool
	}
	Results struct {
		Present bool
		Count   int
	}
	Open struct {
		Present bool
		LinkID  string
	}
}

// ParseParams takes raw commandline arguments and returns a Params structure.
func ParseParams(args []string) (Params, error) {

	var params Params

	//  Flatten the args.
	flatArgs := strings.Join(args, " ")

	//  Handle the help param.
	params.ShowHelp.Present = parseHelpFlag(flatArgs)

	//  Handle the results param.
	var err error
	params.Results.Present, params.Results.Count, err = parseResultsFlag(flatArgs)
	if err != nil {
		return params, err
	}

	//  Handle the open param.
	params.Open.Present, params.Open.LinkID = parseOpenFlag(flatArgs)
	if err != nil {
		return params, err
	}

	return params, nil
}

//	Internally used to check for the help flag.
func parseHelpFlag(flatArgs string) bool {
	match, _ := regexp.MatchString(`-h|--help`, flatArgs)
	return match
}

//	Checks for the results flag and returns its value.
func parseResultsFlag(flatArgs string) (bool, int, error) {
	re := regexp.MustCompile(`(-r|--results)(=?|\s)(?P<ResultsCount>\d+)`)
	matches := re.FindStringSubmatch(flatArgs)
	if len(matches) != 4 {
		return false, 0, nil
	}
	i, err := strconv.Atoi(matches[3])
	return true, i, err
}

//  Given args, returns the value of the --open=linkid param
func parseOpenFlag(flatArgs string) (bool, string) {
	//  todo fix regex for any char
	re := regexp.MustCompile(`(-o|--open)(=?|\s)(?P<LinkId>\d+)`)
	matches := re.FindStringSubmatch(flatArgs)
	if len(matches) != 4 {
		return false, ""
	}
	return true, matches[3]
}
