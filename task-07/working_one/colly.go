package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gocolly/colly"
)

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
	head := [][]string{
		{"Country", "NetWorth", "Age", "Source"},
	}
	for _, head := range head {
		_ = writer.Write(head)
	}

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
	c.OnHTML("tr.base.ng-scope", func(e *colly.HTMLElement) {
		writer.Write([]string{
			e.ChildText("td.Country"),
			e.ChildText("td.Net.Worth"),
			e.ChildText("td.age"),
			e.ChildText("td.source"),
		})
		fmt.Println(e.Text)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	c.Visit("http://localhost:8100")
}
