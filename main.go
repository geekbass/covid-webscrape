package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

/*
Flow:
1) Scrape the Site to retrieve stats
2) Add the stats to DB
3) Display as website
4) serve stats via API
*/

func main() {
	url := "https://coronavirus.ohio.gov/wps/portal/gov/covid-19/"

	// Initialize New Collector
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10.14; rv:73.0) Gecko/20100101 Firefox/73.0"),
	)

	// Get the stats
	c.OnHTML(".odh-ads__item", func(e *colly.HTMLElement) {
		data := e.ChildText("[class=odh-ads__item-title]")
		summary := e.ChildText("[class=odh-ads__item-summary]")
		fmt.Println(data, summary)
	})

	// Get the counties
	c.OnHTML(".odh-ads__container", func(e *colly.HTMLElement) {
		counties := e.ChildText("[class=odh-ads__super-script-item]")
		fmt.Println(counties)
	})

	// Scrape
	c.Visit(url)

}
