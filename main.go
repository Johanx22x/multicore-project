package main

import (
	"fmt"

	"github.com/Johanx22x/multicore-project/internal/scrap"
	// "github.com/Johanx22x/multicore-project/internal/server"
	"github.com/Johanx22x/multicore-project/internal/csv"
)


func main() {
    // scrap.ObtainCSV()
    // server.Start()
    urls := csv.GetURLs()
    for _, url := range urls {
        fmt.Println(url)
        info, err := scrap.GetMetadata(url[1])  
        if err != nil {
            fmt.Println(err)
        } else {
            fmt.Println(string(info))
        }
    }
}
