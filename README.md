<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/Johanx22x/multicore-project/images/golangroutines-2612414795.png">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a>

  <h3 align="center">Websites Content Analysis Using Multithreading</h3>
</div>

Terminal based program to analyze the metadata of the 1000 most visited websites around the world, 
implementing [Go programming language](https://go.dev/), web scraping with [crawlerclub](https://github.com/crawlerclub/ce), 
multithreading using goroutines and chart libraries with [go-echarts](https://github.com/go-echarts/go-echarts). 
<!--and a local web server to host the charts.-->

The information is stored in JSON and CSV files.

[![JSON][json]](https://github.com/Johanx22x/multicore-project/images/gopher-vs-json.png)

# Main Menu

[![Main Menu][main-menu]](https://github.com/Johanx22x/multicore-project/images/main-menu.png)

# Obtained charts example

[![Chart example][chart-example]](https://github.com/Johanx22x/multicore-project/images/chart-example.png)

# Running instructions

## Dependencies

- [go](https://go.dev/dl/)

## Run the program 
```bash
$ go run main.go
```

# Warning

Do not compile the source code, there is a null pointer error if you try to run the binary file.

Instead use `go run main.go` to run the program.
