package main

import (
	"github.com/Johanx22x/multicore-project/internal/jsonm"
	"github.com/Johanx22x/multicore-project/internal/scrap"
)

const WORKERS = 15

func main() {
    scrap.ObtainCSV()
    jsonm.SaveMetadata(WORKERS)
    jsonm.LoadMetadata()
}
