package main

import "github.com/fatih/color"
import "log"
import "fmt"
import "os"
import "regexp"
import "strconv"
import "strings"

func main() {

    //  Get the number of command line parameters.
    args := os.Args[1:]
    argCount := len(args)
    flatArgs := strings.Join(args, " ")

    //  Bail if we have too few arguments.
    if(argCount < 2) {
        showHelp()
    }

    fmt.Println("Args: ", args)

    helpFlag := ParseHelpFlag(flatArgs)
    if helpFlag == true {
        showHelp()
        os.Exit(0)
    }

    r := regexp.MustCompile(`(-r|--results)(=?|\s)(?P<ResultsCount>\d+)`)
    fmt.Println(r)


    resultsCount, err := ParseResultsFlag(flatArgs, 3)
    if err != nil {
        showHelp()
        os.Exit(1)
    }
    query := args[len(args)-1]

    //  Get the environment variables.
    apiKey := os.Getenv("GOOGLEIT_API_KEY")
    engineId := os.Getenv("GOOGLEIT_ENGINE_ID")

    // fmt.Println("Help Flag:     ", helpFlag)
    // fmt.Println("Results Count: ", resultsCount)
    // fmt.Println("Query:         ", query)

    //  Run the search
    results, err := DoSearch(query, apiKey, engineId)
    if err != nil {
        log.Fatal("Search error: ", err)
        os.Exit(1)
    }

    //  Show the results.
    for i := 0; i < len(results.Items) && i < resultsCount; i++ {
        item := results.Items[i]
        PrintResult(item)
    }
}

func PrintResult(item ResultItem) {
    color.Green(item.Title)
    color.White(item.Snippet)
    color.Cyan(item.Link)
    fmt.Println()
}

func ParseHelpFlag(flatArgs string) bool {
    match, _ := regexp.MatchString(`-h|--help`, flatArgs)
    return match
}

func ParseResultsFlag(flatArgs string, defaultValue int) (int, error) {
    re := regexp.MustCompile(`(-r|--results)(=?|\s)(?P<ResultsCount>\d+)`)
    matches := re.FindStringSubmatch(flatArgs)
    if len(matches) != 4 {
        return defaultValue, nil
    } else {
        i, err := strconv.Atoi(matches[3])
        if err == nil && i > 10 {
            fmt.Println("Results are limited to 10...")
            i = 10
        }
        return i, err
    }
}


func showHelp() {
    fmt.Println("usage: google-it [--results] [--help] <query>")
    fmt.Println("  --results, -r: Number of results to show. Default = 3, Max = 10")
    fmt.Println("  --help, -h:    Show help")
    fmt.Println("  <query>:       Text to search for")
    fmt.Println("")
    fmt.Println("examples:")
    fmt.Println("  google-it \"Population of Indonesia\"")
}
