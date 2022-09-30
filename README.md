# Websites Content Analysis Using Multicore

Terminal based program to analyze the metadata of the 1000 most visited websites around the world, 
implementing [Go programming language](https://go.dev/), web scraping with [crawlerclub](https://github.com/crawlerclub/ce), 
multithreading using goroutines, chart libraries with [go-echarts](https://github.com/go-echarts/go-echarts) 
and a small local web server to host the charts.

# Compiling instructions

## Dependencies

- [go](https://go.dev/dl/)

## Run the program 
```bash
$ go run main.go
```

# Warning

Do not compile the source code, there is a null pointer error if you try to run the binary file.

Instead use `go run main.go` to run the program.
