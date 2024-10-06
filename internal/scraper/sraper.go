package scraper

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/tbauriedel/firefighting-operations-nland/internal/operations"
)

const allowedDomains = "www.kfv-online.de"

type Scraper struct {
	RC         int
	Err        error
	Failure    bool
	Operations []operations.Operation
	Collector  *colly.Collector
}

func New() (s *Scraper) {
	s = &Scraper{
		RC:         0,
		Err:        nil,
		Failure:    false,
		Operations: nil,
		Collector:  nil,
	}

	//log.Print("Register new scraper")

	s.Collector = colly.NewCollector(
		colly.AllowedDomains(allowedDomains),
	)

	s.Collector.OnError(func(r *colly.Response, err error) {
		s.Err = fmt.Errorf("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
		s.Failure = true
	})

	s.Collector.OnResponse(func(r *colly.Response) {
		s.RC = r.StatusCode
	})

	//log.Print("New scraper registered")

	return
}

func (s *Scraper) RegisterOperations() {
	s.Collector.OnHTML("table#operationList > tbody", func(HtmlElement *colly.HTMLElement) {
		HtmlElement.ForEach("tr", func(_ int, element *colly.HTMLElement) {

			row := operations.Operation{
				Time:     element.ChildText("td:nth-child(1)"),
				Units:    operations.ProcessUnits(element.ChildText("td:nth-child(2)")),
				District: element.ChildText("td:nth-child(3)"),
				Report:   element.ChildText("td:nth-child(4)"),
				Location: element.ChildText("td:nth-child(5)"),
			}
			s.Operations = append(s.Operations, row)
		})
	})

	//log.Print("Register operations")
}
