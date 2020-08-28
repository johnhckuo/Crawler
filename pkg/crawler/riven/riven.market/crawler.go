package RivenMarket

import (
	"log"

	"github.com/gocolly/colly"
)

type Crawler struct {
}

func (*Crawler) GetRivenByWeapon(targetWeapon *string) (*string, error) {

	domain := "riven.market"
	rootUrl := "https://" + domain + "/_modules/riven/showrivens.php?platform=PC&limit=500&recency=-1&veiled=false&onlinefirst=false&polarity=all&rank=all&mastery=16&stats=Any&neg=all&price=99999&rerolls=-1&sort=time&direction=ASC&page=1&time=1598629713396&weapon="
	targetUrl := rootUrl + *targetWeapon

	var response string

	// Instantiate default collector
	c := colly.NewCollector(
		// Turn on asynchronous requests
		//colly.Async(true),
		// Attach a debugger to the collector
		//colly.Debugger(&debug.LogDebugger{}),
		colly.AllowedDomains(domain),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("#riven-list", func(e *colly.HTMLElement) {
		//t, a := e.DOM.Html()
		//log.Printf("Link found: %+v, %v", t, a)

		e.ForEach(".weapon .xs-none", func(_ int, el *colly.HTMLElement) {
			stats, _ := el.DOM.Html()
			log.Printf("Link found: %v", stats)

		})
		//link := e.Attr("href")
		// Print link
		//log.Printf("Link found: %q -> %s\n", e.Text, link)

		response = e.Text
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		//c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		log.Printf("Retrieving %s from %s", *targetWeapon, targetUrl)
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		log.Printf("Request URL:%v failed with response: %+v \nError:%+v", r.Request.URL, r, err.Error())
	})

	// Start scraping
	c.Visit(targetUrl)

	return &response, nil
}
func (*Crawler) GetRivenByStats(stats []string) (*string, error) {
	rootUrl := "https://riven.market/list/PC"
	targetWeapon := "Acceltra"
	targetUrl := rootUrl + "/Acceltra"

	var response string

	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains(rootUrl),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		log.Printf("Link found: %q -> %s\n", e.Text, link)
		response = e.Text
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		//c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		log.Printf("Retrieving %s from %s", targetWeapon, targetUrl)
	})

	// Start scraping
	c.Visit(targetUrl)

	return &response, nil
}
