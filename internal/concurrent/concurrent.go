package concurrent

// Create chunks from a list of string
func ChunkSlice(urlList []string, WORKERS int) (chunkList [][]string) {
    for i := 0; i < len(urlList); i += WORKERS {
        end := i + WORKERS
        if end > len(urlList) {
            end = len(urlList)
        }
        chunkList = append(chunkList, urlList[i:end])
    }
    // Return chunkList
    return
}
