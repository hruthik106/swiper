package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	titleFlag := flag.Bool("title", false, "shows only the page title")
	linksFlag := flag.Bool("links", false, "shows only the links on the page")

	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("usage :  swiper <url> [--title] [--links]")
		os.Exit(1)
	}

	url := args[0]
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("error creating a request : ", err)
		os.Exit(1)
	}
	req.Header.Set("user-Agent", "swiper/1.0(+https://github.com/hrithik106/swiper)")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error fetching url", url)
		os.Exit(1)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Failed to fetch page . status : ", resp.Status)
		os.Exit(1)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("ERROR parsing HTML :", err)
		os.Exit(1)
	}

	title := doc.Find("title").Text()
	var links []string
	doc.Find("a").Each(func(i int, s *goquery.Selection) {

		href, exists := s.Attr("href")
		if exists {
			fulllink := makeAbsolute(url, href)
			links = append(links, fulllink)
		}
	})
	if *titleFlag {
		fmt.Println("Title : ", title)
		return
	}

	if *linksFlag {
		fmt.Println("Links :")
		for i, link := range links {
			fmt.Printf("%d. %s\n", i+1, link)
		}
		fmt.Printf("\nTotal links found : %d\n ", len(links))
		return
	}

	fmt.Println("title :", title)
	fmt.Println("links found : ", len(links))

	fmt.Println("\n review (first 5 links)")
	limit := 5
	if len(links) < 5 {
		limit = len(links)
	}

	for i := 0; i < limit; i++ {
		fmt.Printf("%d. %s\n", i+1, links[i])
	}

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
