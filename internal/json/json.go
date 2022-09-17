package json

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	// "sync"

	"crawler.club/ce"
	"github.com/Johanx22x/multicore-project/internal/csv"
	"github.com/Johanx22x/multicore-project/internal/scrap"
)

// func SaveMetadata(wg *sync.WaitGroup) {
//     defer wg.Done()
func SaveMetadata() {

    urls := csv.GetURLs()
    var metaList = make(map[string]ce.Doc)

    // TODO: Implement concurrent mode to get data
    for idx, url := range urls {
        if idx == 3 {
            break
        }
        fmt.Println(url)
        pageName := url[1]

        info, _ := scrap.Metadata(pageName)  

        metaList[pageName] = info
    }

    jsonFile, err := json.Marshal(metaList)
    if err != nil {
        log.Fatal(err)
    }

    ioutil.WriteFile("Data/websites-metadata.json", jsonFile, os.ModePerm)
}
