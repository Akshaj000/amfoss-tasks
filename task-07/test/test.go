package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func scrap() {

	fName := "data1.csv"
	file, err := os.Create(fName)
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

func main() {
	scrap()
	csvfile1, err := os.Open("data1.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer csvfile1.Close()
	reader := csv.NewReader(csvfile1)
	csvfile2, err := os.Create("data.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	writer := csv.NewWriter(csvfile2)

	for i := 0; i < 10; i++ {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			return
		}
		err = writer.Write(record)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	writer.Flush()
	err = csvfile2.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	e := os.Remove("data1.csv")
	if e != nil {
		log.Fatal(e)
	}
}
