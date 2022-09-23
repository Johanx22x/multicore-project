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

// Save metadata obtained from the most visited websites
//
// Receive the number of workers to implement in concurrency
func SaveMetadata(WORKERS int) {
    // Create the WaitGroup
    wg := sync.WaitGroup{}

    // Get the stored urls
    idexUrls := csv.GetURLs()

    // Create a list with only the url (exclude the index)
    var urls []string
    for _, val := range idexUrls {
        urls = append(urls, val[1])
    }

    // Create chunk slices with the WORKERS size
    urlList := concurrent.ChunkSlice(urls, WORKERS)

    // Create a buffer to store the channels
    buffers := make([]chan ce.Doc, len(urlList))

    // Iterate over the chunks
    for idx, urls := range urlList {
        // Create a temporal channel
        tmpchan := make(chan ce.Doc, len(urls))

        // For any single url in the chunk
        for _, url := range urls {
            // Add a worker
            wg.Add(1)
            // Get metadata
            go scrap.Metadata(&wg, tmpchan, url)
        }
        // Add the tmpchan to the buffer
        buffers[idx] = tmpchan
    }
    // Wait all the workers
    wg.Wait()

    // Create a map to store the obtained metadata
    var metaList = make(map[string]ce.Doc)

    // Iterate over the chunks
    for i := 0; i < len(urlList); i++ {
        // While the buffer isn't empty
        for len(buffers[i]) > 0 {
            // Add the information to the metadata map
            response := <-buffers[i]
            metaList[response.Url] = response
        }
    
    }

    // Convert the metadata list to a json format 
    jsonFile, err := json.Marshal(metaList)
    if err != nil {
        log.Fatal(err)
    }

    // Save the metadata obtained in json format
    ioutil.WriteFile("Data/websites-metadata.json", jsonFile, os.ModePerm)
}

// Load the obtained data from the JSON file
func LoadMetadata() {
    // Read the file and error management
    content, err := ioutil.ReadFile("Data/websites-metadata.json")
    if err != nil {
        log.Fatal(err)
    }

    // Create a map for store the content
    payload := make(map[string]ce.Doc)
    err = json.Unmarshal(content, &payload)
    if err != nil {
        log.Fatal("Error during Unmarshal(): ", err)
    }

    // TODO: implement data return
    cont := 0
    for _, page := range payload {
        fmt.Println(cont, page.Url, page.Text)
        cont++
    }
}
