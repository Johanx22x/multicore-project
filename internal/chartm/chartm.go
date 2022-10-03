package chartm

import (
	"bufio"
	"log"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

func CreateChart(dataMap map[string]float64, title, name string) {
    newPieData := []opts.PieData{}
    for key, val := range dataMap {
        tmpPie := opts.PieData{Name: key, Value: val}
        newPieData = append(newPieData, tmpPie)
    }

    dataToChart := newPieData
    pie := charts.NewPie()

    pie.SetGlobalOptions(
        charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeChalk}),
        charts.WithTitleOpts(opts.Title{Title: title}),
    )

    pie.AddSeries(name, dataToChart)

    os.Mkdir("./Data/Chart/" + name, os.ModePerm);
	f, _ := os.Create("./Data/Chart/" + name + "/" + name + ".html")
	defer f.Close()

    pie.Render(f)
}

func SaveWords(words []string, name string) {
    path := "./Data/Chart/" + name + "/" + name + "txt"
	f, _ := os.Create(path)
	defer f.Close()

    file, err := os.OpenFile(path , os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }

    totalWords := ""
    for _, word := range words {
        totalWords += word + " "
    }

    textWriter := bufio.NewWriter(file)
    defer textWriter.Flush()

    // write string to buffer before saving file.
    _, err = textWriter.WriteString(totalWords)
    if err != nil {
        log.Fatal(err)
    }
}
