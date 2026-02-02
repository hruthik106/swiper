package scraper

import (
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

const testHTML = `
<!DOCTYPE html>
<html>
<head>
	<title>Test Page</title>
</head>
<body>
	<h1>Hello World</h1>
	<a href="https://example.com">Example</a>
	<a href="/about">About</a>
</body>
</html>
`

func TestScrapeFromHTML(t *testing.T) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(testHTML))
	if err != nil {
		t.Fatal(err)
	}

	// title
	title := doc.Find("title").Text()
	if title != "Test Page" {
		t.Fatalf("expected title 'Test Page', got '%s'", title)
	}

	// links
	var links []string
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		links = append(links, href)
	})

	if len(links) != 2 {
		t.Fatalf("expected 2 links, got %d", len(links))
	}
}

func TestSelectorScraping(t *testing.T) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(testHTML))
	if err != nil {
		t.Fatal(err)
	}

	text := strings.TrimSpace(doc.Find("h1").First().Text())
	if text != "Hello World" {
		t.Fatalf("expected 'Hello World', got '%s'", text)
	}
}
