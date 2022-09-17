package json

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

    "github.com/crawlerclub/ce"
	"github.com/Johanx22x/multicore-project/internal/csv"
	"github.com/Johanx22x/multicore-project/internal/scrap"
)

func SaveMetadata() {
    urls := csv.GetURLs()
    var metaList = make(map[string]ce.Doc)

    // TODO: Implement concurrent mode to get data
    for idx, url := range urls {
        if idx == 2 {
            break
        }
        pageName := url[1]
        fmt.Println(pageName)

        info := scrap.Metadata(pageName)

        metaList[pageName] = info
    }

    jsonFile, err := json.Marshal(metaList)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(jsonFile))

    ioutil.WriteFile("Data/websites-metadata.json", jsonFile, os.ModePerm)
}

func LoadMetadata() {
    content, err := ioutil.ReadFile("Data/websites-metadata.json")
    if err != nil {
        log.Fatal(err)
    }

    payload := make(map[string]ce.Doc)
    err = json.Unmarshal(content, &payload)

    if err != nil {
        log.Fatal("Error during Unmarshal(): ", err)
    }
    for _, page := range payload {
        fmt.Println(page.Title)
        fmt.Println(page.Author)
        fmt.Println(page.CanonicalUrl)
        fmt.Println(page.Favicon)
        fmt.Println(page.Debug)
        fmt.Println(page.From)
        fmt.Println(page.Html)
        fmt.Println(page.Images)
        fmt.Println(page.Language)
        fmt.Println(page.Location)
        fmt.Println(page.Published)
        fmt.Println(page.PublishedParsed)
        fmt.Println(page.Tags)
        fmt.Println(page.Text)
        fmt.Println(page.Url)
    }
}
