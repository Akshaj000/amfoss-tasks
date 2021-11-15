package main

import (
	"encoding/csv"
	"encoding/json"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"github.com/geziyor/geziyor/export"
)

type Application struct {
	Name     string
	Networth string
	Age      string
	Country  string
	Source   string
}

func main() {

	fileone, _ := os.Create("out.json")
	fileone.Close()
	geziyor.NewGeziyor(&geziyor.Options{
		StartRequestsFunc: func(g *geziyor.Geziyor) {
			g.GetRendered("https://www.forbes.com/real-time-billionaires/#50cd30c23d78", g.Opt.ParseFunc)
		},
		ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {
			r.HTMLDoc.Find("tr.base.ng-scope").Each(func(_ int, s *goquery.Selection) {
				g.Exports <- map[string]interface{}{
					"Country":  s.Find("td.Country\\/Territory").Text(),
					"Networth": s.Find("td.Net.Worth").Text(),
					"Age":      s.Find("td.age").Text(),
					"Source":   s.Find("td.source").Text(),
				}
			})
		},
		Exporters: []export.Exporter{&export.JSON{}},
	}).Start()

	body, _ := os.ReadFile("out.json")
	sbody := string(body)
	sbody1 := ""
	count := 0
	for i := 0; i < len(sbody); i++ {
		if count == 10 {
			sbody1 += "]"
			break
		} else {
			sbody1 += string(sbody[i])
		}
		if string(body[i]) == "}" {
			count += 1
		}

	}
	f, _ := os.Create("out.json")
	f.Write([]byte(sbody1))
	f.Close()

	csvFile, _ := os.Create("./data.csv")
	defer csvFile.Close()
	jsonDataFromFile, _ := os.ReadFile("./out.json")
	var jsonData []Application
	json.Unmarshal([]byte(jsonDataFromFile), &jsonData)
	writer := csv.NewWriter(csvFile)
	head := [][]string{
		{"Name", "Net-Worth", "Age", "Country", "Source"},
	}
	for _, head := range head {
		_ = writer.Write(head)
	}
	for _, usance := range jsonData {
		var row []string
		row = append(row, usance.Networth)
		row = append(row, usance.Age)
		row = append(row, usance.Country)
		row = append(row, usance.Source)
		writer.Write(row)
	}
	writer.Flush()
	os.Remove("out.json")

}
