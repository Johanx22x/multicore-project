package main

import (
    "github.com/Johanx22x/multicore-project/internal/scrap"
    "github.com/Johanx22x/multicore-project/internal/server"
)

func main() {
    scrap.ObtainCSV()
    server.Start()
}
