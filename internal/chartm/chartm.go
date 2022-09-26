package chartm

import (
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

	f, _ := os.Create("./Data/Chart/" + name + ".html")
	defer f.Close()

    pie.Render(f)
}

