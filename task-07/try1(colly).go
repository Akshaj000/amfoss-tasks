package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {

	file, err := os.Create("data.csv")
	if err != nil {
		log.Fatalf("Could not create file, err :%q", err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	c := colly.NewCollector()
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error:", err.Error())
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("visited", r.Request.URL)
	})
	var a = 0
	c.OnHTML("h2.winners-losers-headline", func(e *colly.HTMLElement) {
		for a < 10 {
			writer.Write([]string{
				e.ChildText("td.Net.Worth"),
				e.ChildText("td.age"),
				e.ChildText("td.Country/Terrirory"),
				e.ChildText("td.source"),
			})
		}
		e.Request.Visit(e.Attr("tr.base.ng-scope"))
		fmt.Println(e.Text)
	})
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})
	c.Visit("https://www.forbes.com/real-time-billionaires/#56138aab3d78")

}
