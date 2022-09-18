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

func Metadata(wg *sync.WaitGroup, out chan ce.Doc, url string) {
    defer wg.Done()

    url = fmt.Sprintf("https://%s", url)

	res := dl.DownloadUrl(url)

	items := strings.Split(res.RemoteAddr, ":")
	ip := ""
	if len(items) > 0 {
		ip = items[0]
	}
	doc := *ce.ParsePro(url, res.Text, ip, true)

    out <- doc
}

func ObtainCSV() {
    os.Mkdir("Data", os.ModePerm)

    fName := "Data/top-websites.csv"
    file, _ := os.Create(fName)
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
