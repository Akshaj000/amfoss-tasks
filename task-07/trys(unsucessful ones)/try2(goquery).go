package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func Parseurl(url string) []string {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	doc, _ := goquery.NewDocumentFromReader(resp.Body)
	var data []string

	doc.Find("p").Each(func(_ int, s *goquery.Selection) {
		s.Find("#text").Each(func(_ int, t *goquery.Selection) {
			datas := t.Text()
			data = append(data, datas)
		})
	})
	return data
}

func main() {
	url := "https://www.forbes.com/real-time-billionaires/#5eb61b473d78"

	// Output filename
	out_filename := "data.csv"
	data := Parseurl(url)
	file, _ := os.Create(out_filename)
	defer file.Close()
	var err error
	for _, datas := range data {
		_, err = io.WriteString(file, datas)

		if err != nil {
			log.Fatal(err)
		}

		file.Sync()
	}
}
