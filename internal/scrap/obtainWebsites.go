package scrap

import (
    "os"
    "encoding/csv"
    "encoding/json"
	"log"
	"strings"
	"fmt"

	"github.com/crawlerclub/ce"
	"github.com/crawlerclub/dl"
	"github.com/gocolly/colly"
)

// func GetQuery(url string) (string, error) {
//     formatUrl := fmt.Sprintf("https://%s", url)

//     fmt.Printf("Scanning %s\n", formatUrl)
// 	resp, err := http.Get(formatUrl)
// 	// handle the error if there is one
// 	if err != nil {
//         return "", err
// 	}
// 	// do this now so it won't be forgotten
// 	defer resp.Body.Close()
// 	// reads html as a slice of bytes
// 	html, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
//         return "", err
// 	}

//     p := strings.NewReader(string(html))
//     doc, _ := goquery.NewDocumentFromReader(p)

//     doc.Find("script").Each(func(i int, el *goquery.Selection) {
//         el.Remove()
//     })

//     return doc.Text(), nil
// }

func GetMetadata(url string) ([]byte, error) {
    url = fmt.Sprintf("https://%s", url)

	res := dl.DownloadUrl(url)
	if res.Error != nil {
        return nil, res.Error
	}

	items := strings.Split(res.RemoteAddr, ":")
	ip := ""
	if len(items) > 0 {
		ip = items[0]
	}
	doc := ce.ParsePro(url, res.Text, ip, true)
    j, _ := json.Marshal(doc)
    return j, nil
}

func ObtainCSV() {
    os.Mkdir("Data", os.ModePerm)

    fName := "Data/top-websites.csv"
    file, err := os.Create(fName)
    if err != nil {
        log.Fatalf("Could not create file, err: %q", err)
        return
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    c := colly.NewCollector(
        colly.AllowedDomains("www.htmlstrip.com"),
    )

    // Find and print all links
    c.OnHTML(".row", func(e *colly.HTMLElement) {
        row := e.ChildText(".col-6")

        idAndValue := strings.Fields(row)

        if len(idAndValue) > 1 {
            id := idAndValue[0]
            value := idAndValue[1]

            writer.Write([]string{
                id,
                value,
            })
        }
    })

    c.Visit("https://www.htmlstrip.com/alexa-top-1000-most-visited-websites#")
}

