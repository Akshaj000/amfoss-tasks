package main

import (
	"encoding/csv"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
)

func main() {

	file, _ := os.Create("data.csv")
	writer := csv.NewWriter(file)
	head := []string{"Country", "Networth", "Age", "Source"}
	writer.Write(head)
	count := 0
	geziyor.NewGeziyor(&geziyor.Options{
		StartRequestsFunc: func(g *geziyor.Geziyor) {
			g.GetRendered("https://www.forbes.com/real-time-billionaires/#50cd30c23d78", g.Opt.ParseFunc)
		},
		ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {
			r.HTMLDoc.Find("tr.base.ng-scope").Each(func(_ int, s *goquery.Selection) {
				Country := s.Find("td.Country\\/Territory").Text()
				Networth := s.Find("td.Net.Worth").Text()
				Age := s.Find("td.age").Text()
				Source := s.Find("td.source").Text()
				posts := []string{Country, Networth, Age, Source}
				writer.Write(posts)
				count += 1
				if count <= 10 {
					writer.Flush()
				}
			})

		},
	}).Start()
}
