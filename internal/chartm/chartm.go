package chartm

import (
	"bufio"
	"log"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

// Create a new chart accordng to a map parameter
func CreateChart(dataMap map[string]float64, title, name string) {
    // Create a new pie data array
    newPieData := []opts.PieData{}

    // Set the keys and values
    for key, val := range dataMap {
        tmpPie := opts.PieData{Name: key, Value: val}
        newPieData = append(newPieData, tmpPie)
    }

    // Create a new pie chart struct
    dataToChart := newPieData
    pie := charts.NewPie()
    
    // Set some options for the chart render
    pie.SetGlobalOptions(
        charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeChalk}),
        charts.WithTitleOpts(opts.Title{Title: title}),
    )

    // Add the series to the pie chart
    pie.AddSeries(name, dataToChart)

    // Create the folder to store the chart
    os.Mkdir("./Data/Chart/" + name, os.ModePerm);
	f, _ := os.Create("./Data/Chart/" + name + "/" + name + ".html")
	defer f.Close()

    // Render the pie
    pie.Render(f)
}

// Save the found words of the webpage in a txt file
func SaveWords(words []string, name string) {
    path := "./Data/Chart/" + name + "/" + name + ".txt"

    //Create a new file with the above path
	f, _ := os.Create(path)
	defer f.Close()

    // Open the file
    file, err := os.OpenFile(path , os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }

    // Create a new string with the words
    totalWords := ""
    for _, word := range words {
        totalWords += word + " "
    }

    // Pass string into a buffer writer
    textWriter := bufio.NewWriter(file)
    defer textWriter.Flush()

    // write string to buffer before saving file.
    _, err = textWriter.WriteString(totalWords)
    if err != nil {
        log.Fatal(err)
    }
}
