package main

import (
	"flag"
	"fmt"
	"net/url"
	"strings"

	"github.com/ernesto-jimenez/scraperboard"
)

func main() {
	flag.Parse()

	query := strings.Join(flag.Args(), " ")
	searchURL := fmt.Sprintf("https://www.google.com/search?q=%s", url.QueryEscape(query))

	scraper, _ := scraperboard.NewScraperFromString(scraperXML)

	var response Response
	scraper.ExtractFromURL(searchURL, &response)

	for _, result := range response.Results {
		fmt.Printf("%s:\n\t%s\n", result.Title, result.URL)
	}
}

// Response contains an array of google results (Result)
type Response struct {
	Results []Result
}

// Result has a Title and URL
type Result struct {
	Title string
	URL   string
}

var scraperXML = `
	<Scraper>
		<Each name="results" selector="#search ol li.g">
			<Property name="title" selector="a"/>
			<Property name="url" selector="a">
				<Filter type="first"/>
				<Filter type="attr" argument="href"/>
			</Property>
		</Each>
	</Scraper>
`
