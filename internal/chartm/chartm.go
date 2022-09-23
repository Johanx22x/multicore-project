package chartm

import (
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

func CreateChart(langs map[string]float64, title, name string) {
    langsPie := []opts.PieData{}
    for key, val := range langs {
        tmpPie := opts.PieData{Name: key, Value: val}
        langsPie = append(langsPie, tmpPie)
    }

    languages := langsPie
    pie := charts.NewPie()

    pie.SetGlobalOptions(
        charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeChalk}),
        charts.WithTitleOpts(opts.Title{Title: title}),
    )

    pie.AddSeries(name, languages)

	f, _ := os.Create("./Data/Chart/" + name + ".html")
	defer f.Close()

    pie.Render(f)
}

