package main

import (
	"github.com/Johanx22x/multicore-project/internal/json"
	"github.com/Johanx22x/multicore-project/internal/scrap"
)

const WORKERS = 8

func main() {
    scrap.ObtainCSV()
    json.SaveMetadata()
    json.LoadMetadata()
}
