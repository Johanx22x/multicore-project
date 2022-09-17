package csv

import (
    "os"
    "fmt"
    "encoding/csv"
)

func GetURLs() [][]string {
    file, err := os.Open("Data/top-websites.csv")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()

    reader := csv.NewReader(file)
    records, _ := reader.ReadAll()
 
    return records
}
