package jsonm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"

	"github.com/Johanx22x/multicore-project/internal/concurrent"
	"github.com/Johanx22x/multicore-project/internal/csv"
	"github.com/Johanx22x/multicore-project/internal/scrap"
	"github.com/crawlerclub/ce"
)

func SaveMetadata(WORKERS int) {
    wg := sync.WaitGroup{}

    idexUrls := csv.GetURLs()
    var urls []string
    for _, val := range idexUrls {
        urls = append(urls, val[1])
    }

    urlList := concurrent.ChunkSlice(urls, WORKERS)

    buffers := make([]chan ce.Doc, len(urlList))

    for idx, urls := range urlList {
        tmpchan := make(chan ce.Doc, len(urls))
        for _, url := range urls {
            wg.Add(1)
            go scrap.Metadata(&wg, tmpchan, url)
        }
        buffers[idx] = tmpchan
    }
    wg.Wait()

    var metaList = make(map[string]ce.Doc)

    for i := 0; i < len(urlList); i++ {
        for len(buffers[i]) > 0 {
            response := <-buffers[i]
            metaList[response.Url] = response
        }
    
    }

    jsonFile, err := json.Marshal(metaList)
    if err != nil {
        log.Fatal(err)
    }

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

    cont := 0
    for _, page := range payload {
        fmt.Println(cont, page.Url)
        cont++
    }
}
