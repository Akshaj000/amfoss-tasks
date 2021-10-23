package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {

	fName := "data.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Could not create file, err :%q", err)
		return
	}
	writer := csv.NewWriter(file)
	defer writer.Flush()

	c := colly.NewCollector()
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error:", err.Error())
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("visited", r.Request.URL)
	})
	c.OnHTML("div.table-row", func(e *colly.HTMLElement) {
		writer.Write([]string{
			e.ChildText("div.table-cell.t-name"),
			e.ChildText("div.table-cell.active.t-nw"),
			e.ChildText("div.table-cell.t-country"),
			e.ChildText("div.table-cell.t-industry"),
		})
		fmt.Println(e.Text)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	c.Visit("https://www.bloomberg.com/billionaires/")

}
