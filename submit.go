package main

import (
	"log"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/jroimartin/gocui"
	"github.com/oriiyx/text-voyager/parser"
)

type ResultNavigationData struct {
	title string
	url   string
}

func (a *App) SubmitRequest(g *gocui.Gui, v *gocui.View) error {
	log.Println("submitting")
	popup(g, "Sending request..")
	var googleParser *parser.GoogleResponseParser = &parser.GoogleResponseParser{}

	var r *Request = &Request{
		GoogleParser: *googleParser,
	}

	g.SetCurrentView(SearchPromptView)
	r.SearchQuery = getViewValue(g, SearchPromptView)

	if r.SearchQuery == "" {
		return nil
	}

	log.Println(&r.SearchQuery)
	c := colly.NewCollector(
		// Set the User-Agent to mimic a real browser
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"),
	)

	// Set a random delay to avoid detection
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		RandomDelay: 5 * time.Second,
	})

	c.OnHTML("a[href] h3", func(e *colly.HTMLElement) {
		hrefElement := e.DOM.Parent()
		href, exists := hrefElement.Attr("href")
		if exists {
			resultData := ResultNavigationData{
				title: e.Text,
				url:   href,
			}
			r.ResultNavigationData = append(r.ResultNavigationData, resultData)
		}
	})

	c.OnHTML("#rso", func(e *colly.HTMLElement) {
		// Get the HTML of the #rso element
		var errorHtml error
		r.GoogleParser.SearchElementString, errorHtml = e.DOM.Html()
		// rsoHtml, errorHtml := e.DOM.Html()
		if errorHtml != nil {
			log.Fatalf("Failed to fetch the DOM HTML: %v", errorHtml)
		}

		r.GoogleParser.ParseRawElementString()

		log.Println(r.GoogleParser.ParsedString)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL.String())
	})

	// Handle response
	c.OnResponse(func(r *colly.Response) {
		log.Println("Visited", r.Request.URL)
	})

	// Start scraping on the target URL
	c.Visit("https://www.google.com/search?q=" + r.SearchQuery + "&hl=en&gl=sl&num=15")

	g.CurrentView().Clear()
	g.CurrentView().SetCursor(0, 0)

	// log.Printf("%#v", r.ResultNavigationData)

	return nil
}

func getViewValue(g *gocui.Gui, name string) string {
	v, err := g.View(name)
	if err != nil {
		return ""
	}
	return strings.TrimSpace(v.Buffer())
}
