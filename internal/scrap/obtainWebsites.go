package scrap

import (
    "os"
    "encoding/csv"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

func ObtainCSV() {
    if err := os.Mkdir("Data", os.ModePerm); err != nil {
        log.Fatal(err)
    }

    fName := "Data/top-websites.csv"
    file, err := os.Create(fName)
    if err != nil {
        log.Fatalf("Could not create file, err: %q", err)
        return
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    c := colly.NewCollector(
        colly.AllowedDomains("www.htmlstrip.com"),
    )

    // Find and print all links
    c.OnHTML(".row", func(e *colly.HTMLElement) {
        row := e.ChildText(".col-6")

        idAndValue := strings.Fields(row)

        if len(idAndValue) > 1 {
            id := idAndValue[0]
            value := idAndValue[1]

            writer.Write([]string{
                id,
                value,
            })
        }
    })

    c.Visit("https://www.htmlstrip.com/alexa-top-1000-most-visited-websites#")
}

