package rivenmarket

import (
	"log"
	"strings"

	"github.com/gocolly/colly"
	warframe "github.com/johnhckuo/Crawler/pkg/crawler/warframe"
)

// Crawler struct
type Crawler struct {
	rivens []*warframe.Riven
}

// GetRivenByWeapon retrieve riven by speficied weapon name
func (obj *Crawler) GetRivenByWeapon(targetWeapon *string) (*string, error) {

	domain := "riven.market"
	fetchNumber := "50"
	rootURL := "https://" + domain + "/_modules/riven/showrivens.php?platform=PC&limit=" + fetchNumber + "&recency=-1&veiled=false&onlinefirst=false&polarity=all&rank=all&mastery=16&stats=Any&neg=all&price=99999&rerolls=-1&sort=time&direction=ASC&page=1&time=1598629713396&weapon="
	targetURL := rootURL + *targetWeapon

	// Instantiate default collector
	c := colly.NewCollector(
		// Turn on asynchronous requests
		//colly.Async(true),
		// Attach a debugger to the collector
		//colly.Debugger(&debug.LogDebugger{}),
		colly.AllowedDomains(domain),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("#riven-list", func(rivens *colly.HTMLElement) {
		//t, a := e.DOM.Html()
		//log.Printf("Link found: %+v, %v", t, a)

		rivens.ForEach(".riven", func(_ int, riven *colly.HTMLElement) {

			newRiven := &warframe.Riven{}
			riven.ForEach(".weapon .xs-none", func(_ int, name *colly.HTMLElement) {
				rivenName, _ := name.DOM.Html()
				newRiven.Name = &rivenName
			})
			newRiven.Positives = make(map[string]string)
			riven.ForEach(".pos", func(_ int, positiveElems *colly.HTMLElement) {

				name := strings.TrimSpace(strings.Replace(strings.Replace(positiveElems.DOM.Find(".name").Text(), "\n", "", -1), "Test", "", -1))

				val, _ := positiveElems.DOM.Find(".value input").Attr("value")

				newRiven.Positives[name] = val
			})

			negativeName := strings.TrimSpace(strings.Replace(riven.DOM.Find(".neg").Text(), "\n", "", -1))
			negativeVal, _ := riven.DOM.Find(".neg").Find(".value input").Attr("value")
			newRiven.Negative = make(map[string]string)
			newRiven.Negative[negativeName] = negativeVal

			price, _ := riven.DOM.Find(".price").Find("input").Attr("value")
			newRiven.Price = &price
			seller := strings.TrimSpace(riven.DOM.Find(".seller").Text())
			newRiven.Seller = &seller
			obj.rivens = append(obj.rivens, newRiven)
		})

		for _, riven := range obj.rivens {
			log.Printf("Riven %v , positive %+v, negative %+v, price %v, seller %v", *riven.Name, riven.Positives, riven.Negative, *riven.Price, *riven.Seller)
		}
		//link := e.Attr("href")
		// Print link
		//log.Printf("Link found: %q -> %s\n", e.Text, link)

		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		//c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		log.Printf("Retrieving %s from %s", *targetWeapon, targetURL)
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		log.Fatalf("Request URL:%v failed with response: %+v \nError:%+v", r.Request.URL, r, err.Error())
	})

	// Start scraping
	c.Visit(targetURL)

	return nil, nil
}

// GetRivenByStats retrieve riven by speficied stats
func (*Crawler) GetRivenByStats(stats []string) (*string, error) {
	rootURL := "https://riven.market/list/PC"
	targetWeapon := "Acceltra"
	targetURL := rootURL + "/Acceltra"

	var response string

	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains(rootURL),
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
		log.Printf("Retrieving %s from %s", targetWeapon, targetURL)
	})

	// Start scraping
	c.Visit(targetURL)

	return &response, nil
}
