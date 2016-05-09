package main

import(
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "net/url"
)

func DoSearch(query string, apiKey string, searchEngineId string) (results GoogleResults, err error) {

    //  Build the URL.
    root := `https://www.googleapis.com/customsearch/v1`
    url := fmt.Sprintf("%s?key=%s&cx=%s&q=%s", 
        root, 
        url.QueryEscape(apiKey), 
        url.QueryEscape(searchEngineId),
        url.QueryEscape(query))

    // Build the request
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        log.Fatal("NewRequest: ", err)
        return
    }

    // For control over HTTP client headers,
    // redirect policy, and other settings,
    // create a Client
    // A Client is an HTTP client
    client := &http.Client{}

    // Send the request via a client
    // Do sends an HTTP request and
    // returns an HTTP response
    resp, err := client.Do(req)
    if err != nil {
        log.Fatal("Do: ", err)
        return
    }

    // Callers should close resp.Body
    // when done reading from it
    // Defer the closing of the body
    defer resp.Body.Close()

    // contents, err := ioutil.ReadAll(resp.Body)
    // if err != nil {
    //     log.Fatal("Read: ", err)
    //     return
    // }
    // json := string(contents)

    // fmt.Println("Request")
    // fmt.Println(json)

    // Fill the record with the data from the JSON
    var record GoogleResults

    // Use json.Decode for reading streams of JSON data
    if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
        log.Println(err)
        return record, err
    }

    return record, nil
}
