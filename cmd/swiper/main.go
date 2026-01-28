package main

import (
	"flag"
	"fmt"
	"os"
	"swiper/internal/fileformatter"
	"swiper/internal/scraper"
)

func main() {
	titleFlag := flag.Bool("title", false, "shows only the page title")
	linksFlag := flag.Bool("links", false, "shows only the links on the page")
	jsonFlag := flag.Bool("json", false, "output result as json")
	csvFlag := flag.Bool("csv", false, "output result as csv")
	selectorFlag := flag.String("selector", "", "Scrape elements by css selector")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("usage :  swiper  [--title] [--links] <url>")
		os.Exit(1)
	}

	url := args[0]

	if *selectorFlag != "" {
		results, err := scraper.ScrapeBySelector(url, *selectorFlag)
		if err != nil {
			fmt.Println("error scraping by selector :", err)
			os.Exit(1)
		}

		if *jsonFlag {
			err := fileformatter.OutputAsJSON(map[string]any{
				"selector": *selectorFlag,
				"count":    len(results),
				"results":  results,
			})
			if err != nil {
				fmt.Println("error outputting as json :", err)
				os.Exit(1)
			}
			return
		}

		if *csvFlag {
			err := fileformatter.OutputAsCSV(results)
			if err != nil {
				fmt.Println("error outputting as csv :", err)
				os.Exit(1)
			}
			return
		}

		fmt.Printf("Result for selector \"%s\":\n\n", *selectorFlag)
		for i, r := range results {
			fmt.Printf("%d. %s\n", i+1, r)
		}

		fmt.Printf("\nTotal matches : %d\n", len(results))
		return
	}

	result, err := scraper.Scrape(url)
	if err != nil {
		fmt.Println("error scraping the url :", err)
		os.Exit(1)
	}

	if *titleFlag {
		fmt.Println("Title : ", result.Title)
		return
	}

	if *linksFlag {

		if *jsonFlag {
			err := fileformatter.OutputAsJSON(map[string]any{
				"selector": url,
				"count":    len(result.Links),
				"results":  result.Links,
			})
			if err != nil {
				fmt.Println("error outputting as json :", err)
				os.Exit(1)
			}
			return
		}

		if *csvFlag {
			err := fileformatter.OutputAsCSV(result.Links)
			if err != nil {
				fmt.Println("error outputting as csv :", err)
				os.Exit(1)
			}
			return
		}
		fmt.Println("Links :")
		for i, link := range result.Links {
			fmt.Printf("%d. %s\n", i+1, link)
		}
		fmt.Printf("\nTotal links found : %d\n ", len(result.Links))
		return
	}

	fmt.Println("title :", result.Title)
	fmt.Println("links found : ", len(result.Links))

	fmt.Println("\n review (first 5 links)")
	limit := 5
	if len(result.Links) < 5 {
		limit = len(result.Links)
	}

	for i := 0; i < limit; i++ {
		fmt.Printf("%d. %s\n", i+1, result.Links[i])
	}

}
