package multicore

import (
    "sync"
)

func main() {
    wg := sync.WaitGroup{}
    wg.Add(1)
}
