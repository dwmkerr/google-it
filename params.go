package main

import (
    "regexp"
    "strconv"
    "strings"
)

type Params struct{
    ShowHelp struct{
        Present bool
    }
    Results struct{
        Present bool
        Count int
    }
    Open struct {
        Present bool
        LinkId string
    }
}

func ParseParams(args []string) (Params, error) {

    var params Params

    //  Flatten the args.
    flatArgs := strings.Join(args, " ")

    //  Handle the help param.
    params.ShowHelp.Present = ParseHelpFlag(flatArgs)

    //  Handle the results param.
    var err error
    params.Results.Present, params.Results.Count, err = ParseResultsFlag(flatArgs)
    if err != nil {
        return params, err
    }

    //  Handle the open param.
    params.Open.Present, params.Open.LinkId = ParseOpenFlag(flatArgs)
    if err != nil {
        return params, err
    }

    return params, nil
}

func ParseHelpFlag(flatArgs string) bool {
    match, _ := regexp.MatchString(`-h|--help`, flatArgs)
    return match
}

func ParseResultsFlag(flatArgs string) (bool, int, error) {
    re := regexp.MustCompile(`(-r|--results)(=?|\s)(?P<ResultsCount>\d+)`)
    matches := re.FindStringSubmatch(flatArgs)
    if len(matches) != 4 {
        return false, 0, nil
    } else {
        i, err := strconv.Atoi(matches[3])
        return true, i, err
    }
}

//  Given args, returns the value of the --open=linkid param
func ParseOpenFlag(flatArgs string) (bool, string) {
    //  todo fix regex for any char
    re := regexp.MustCompile(`(-o|--open)(=?|\s)(?P<LinkId>\d+)`)
    matches := re.FindStringSubmatch(flatArgs)
    if len(matches) != 4 {
        return false, ""
    } else {
        return true, matches[3]
    }
}
