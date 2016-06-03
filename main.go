package main

import "github.com/fatih/color"
import "log"
import "fmt"
import "os"
import "strconv"
import "time"
import "github.com/pkg/browser"

const throttleLimit = 100 // no more than 100 API calls per day
const maxResultsCount = 10
const defaultResultsCount = 3

func main() {

	//  Get the number of command line parameters.
	args := os.Args
	argCount := len(args)

	//  Read the parameters.
	params, err := ParseParams(args)
	if err != nil {
		fmt.Println("Error parsing params: ", err)
		os.Exit(1)
	}

	//  Show help and terminate if we have zero params or the help param.
	if argCount == 0 || params.ShowHelp.Present {
		showHelp()
		os.Exit(0)
	}

	//  Handle the results.
	var resultsCount = defaultResultsCount
	if params.Results.Present {
		resultsCount = params.Results.Count
	}
	if resultsCount > 10 {
		fmt.Println("Results are limited to 10...")
		resultsCount = 10
	}

	query := args[len(args)-1]

	//  Get the environment variables.
	apiKey := os.Getenv("GOOGLEIT_API_KEY")
	engineID := os.Getenv("GOOGLEIT_ENGINE_ID")

	//	If there are no API variables, bail.
	if apiKey == "" || engineID == "" {
		color.Red("Missing GOOGLEIT_API_KEY or GOOGLEIT_ENGINE_ID environment variables!")
		color.White("Check the documentation on github.com/dwmkerr/google-it for instructions.")
		os.Exit(1)
	}

	//  Load the settings.
	settings, err := LoadSettings()
	if err != nil {
		color.Red("Error loading settings: ", err)
		os.Exit(1)
	}

	openLinkID := params.Open.LinkID
	if openLinkID != "" {
		for i := 0; i < len(settings.Links); i++ {
			if settings.Links[i].ID == openLinkID {
				color.Green("Opening [%s] %s...", openLinkID, settings.Links[i].URI)
				browser.OpenURL(settings.Links[i].URI)
				os.Exit(0)
			}
		}
		color.Red("Couldn't find link %s to open...", openLinkID)
		os.Exit(0)
	}

	//  Save the settings.
	SaveSettings(settings)

	//  If our throttle settings are for an old date, then
	//  we can reset the API counter.
	today := time.Now()
	if areDatesEqual(today, settings.Throttling.Today) == false {
		//	We are on a new day, reset the throttling.
		settings.Throttling.Calls = 0
		settings.Throttling.Today = today
	} else {
		if settings.Throttling.Calls > throttleLimit {
			color.Red("%d calls made already today, quitting.", settings.Throttling.Calls)
			os.Exit(1)
		}
	}

	// fmt.Println("Help Flag:     ", helpFlag)
	// fmt.Println("Results Count: ", resultsCount)
	// fmt.Println("Query:         ", query)

	//  Run the search
	results, err := DoSearch(query, apiKey, engineID)
	if err != nil {
		log.Fatal("Search error: ", err)
		os.Exit(1)
	}

	//  Show the results, storing any links we show in the settings.
	fmt.Println()
	settings.Links = printResults(results.Items, resultsCount)

	//  Update our throttle settings.
	settings.Throttling.Calls++
	color.White("%d/%d searches made today...", settings.Throttling.Calls, throttleLimit)

	//  Save any updates to our settings.
	err = SaveSettings(settings)
	if err != nil {
		fmt.Println("Error writing settings: ", err)
		color.Red("Warning: When settings are not written, API call counts are not recorded. More than 100 calls per-day will lead to a credit card charge.")
		os.Exit(1)
	}
}

func areDatesEqual(time1 time.Time, time2 time.Time) bool {
	return time1.Year() == time2.Year() && time1.YearDay() == time2.YearDay()
}

func printResults(items []ResultItem, maxResults int) []Link {
	l := 0
	if len(items) > maxResults {
		l = maxResults
	} else {
		l = len(items)
	}
	links := make([]Link, l)
	for i := 0; i < len(items) && i < maxResults; i++ {
		links[i] = printResult(items[i], i)
	}
	return links
}

func printResult(item ResultItem, linkNumber int) Link {
	color.Green(item.Title)
	color.White(item.Snippet)
	color.Cyan("[%d] %s", linkNumber, item.Link)
	fmt.Println()
	return Link{
		ID:  strconv.Itoa(linkNumber),
		URI: item.Link,
	}
}

func showHelp() {
	fmt.Println("usage: google-it [--results 2] [--open 4] [--help] <query>")
	fmt.Println("  --open, -o    : Open link id")
	fmt.Println("  --results, -r : Number of results to show. Default = 3, Max = 10")
	fmt.Println("  --help, -h    : Show help")
	fmt.Println("  <query>       : Text to search for")
}
