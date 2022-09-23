package scrap

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/crawlerclub/ce"
	"github.com/crawlerclub/dl"
	"github.com/gocolly/colly"
)

// Get metadata from a website
//
// Metadata receives a WaitGroup, a channel to send the result and the url of the website
// Metadata sends the result to the channel, doesn't return any data
func Metadata(wg *sync.WaitGroup, out chan ce.Doc, url string) {
    defer wg.Done()

    // Format the url to a valid format
    url = fmt.Sprintf("https://%s", url) 

    // Response from the website
	res := dl.DownloadUrl(url)

    // Obtain items from the response
	items := strings.Split(res.RemoteAddr, ":")
	ip := ""
	if len(items) > 0 {
		ip = items[0]
	}

    // Parse the content
	doc := *ce.ParsePro(url, res.Text, ip, true)

    // Send the obtained data to the channel
    out <- doc
}

// Obtain a CSV table with the most viewed websites in the world
// 
// Store the data in "./Data/top-websites.csv"
func ObtainCSV() {
    // Make a directory to store the data
    os.Mkdir("Data", os.ModePerm)

    // New file name
    fName := "Data/top-websites.csv"
    // Create the file
    file, _ := os.Create(fName)
    // Close the file at the end of the function
    defer file.Close()

    // Create a new csv writer
    writer := csv.NewWriter(file)
    // Flush the writer at the end of the function
    defer writer.Flush()

    // New colly Collector
    c := colly.NewCollector(
        // Only search in allowed domains
        colly.AllowedDomains("www.htmlstrip.com"),
    )

    // Search .row component in the page
    c.OnHTML(".row", func(e *colly.HTMLElement) {
        // Get the rows
        row := e.ChildText(".col-6")

        // Format the row
        idAndValue := strings.Fields(row)

        // Exclude the first row
        if len(idAndValue) > 1 {
            id := idAndValue[0]
            value := idAndValue[1]

            // Write all the allowed rows in the file
            writer.Write([]string{
                id,
                value,
            })
        }
    })

    // Visit the website with the data table 
    c.Visit("https://www.htmlstrip.com/alexa-top-1000-most-visited-websites#")
}
