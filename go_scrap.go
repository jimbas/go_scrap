package main

import (
	"fmt"
	"os"
	"encoding/csv"
	"strings"
	"github.com/gocolly/colly"
)

var categName string = "None"

func main() {

	file, err := os.Create("result.csv")
	if err != nil {
		fmt.Println("Cannot create file:", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Comma = '\t'
	defer writer.Flush()

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	c.OnHTML("div[class=prod-name]", func(e *colly.HTMLElement) {
		name := strings.Replace(e.Text, "  ", "", -1)
		prod := strings.Split(name, "\n")
		for i:=0; i<len(prod); i++ {
			if prod[i] != "" && prod[i] != "New" && prod[i] != "New " {
				fmt.Printf("  Product name:%d %s\n", i, prod[i])
				row := make([]string, 2)
				row[0] = categName
				row[1] = prod[i]
				err := writer.Write(row)
				if err != nil {
					fmt.Println("Cannot write to file:", err)
				}
			}
		}
	})
	c.OnHTML("div[class=mb-24]", func(e *colly.HTMLElement) {
		categUrl := e.ChildAttr("a", "href")
		categName = e.ChildText("p")
		fmt.Printf("Kategory: %s\n", categName)
    	e.Request.Visit(categUrl)
	})
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})
	c.Visit("https://www.orami.co.id/")
}
