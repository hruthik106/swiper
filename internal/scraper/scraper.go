package scraper

import (
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Result struct {
	Title string
	Links []string
}

func Scrape(url string) (*Result, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Swiper/1.0")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	result := &Result{}

	// Title
	result.Title = doc.Find("title").Text()

	// Links
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if ok {
			result.Links = append(result.Links, makeAbsolute(url, href))
		}
	})

	return result, nil
}

func makeAbsolute(base, link string) string {
	if strings.HasPrefix(link, "http") {
		return link
	}
	if strings.HasPrefix(link, "/") {
		return base + link
	}
	return base + "/" + link
}
