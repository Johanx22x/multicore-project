package csv

import (
    "os"
    "fmt"
    "encoding/csv"
)

// Get the urls from the csv file
func GetURLs() [][]string {
    // Open the file
    file, err := os.Open("Data/top-websites.csv")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()

    // Create the new reader
    reader := csv.NewReader(file)
    // Read all the file
    records, _ := reader.ReadAll()
 
    return records
}
