package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/hruthik106/swiper/internal/fileformatter"
	"github.com/hruthik106/swiper/internal/scraper"
	"github.com/hruthik106/swiper/internal/version"
)

func main() {
	var titleFlag bool
	flag.BoolVar(&titleFlag, "title", false, "shows only the page title")
	flag.BoolVar(&titleFlag, "t", false, "shows only the page title")

	var linksFlag bool
	flag.BoolVar(&linksFlag, "links", false, "shows only the links on the page")
	flag.BoolVar(&linksFlag, "l", false, "shows only the links on the page")

	jsonFlag := flag.Bool("json", false, "output result as json")
	csvFlag := flag.Bool("csv", false, "output result as csv")
	selectorFlag := flag.String("selector", "", "Scrape elements by css selector")
	outfile := flag.String("out", "", "output the result into a file")
	versonflag := flag.Bool("version", false, "shows the version of swiper")

	flag.Parse()
	writer, cleanup, err := getWriter(*outfile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Output error:", err)
		os.Exit(1)
	}

	args := flag.Args()
	if len(args) < 1 {
		Cli_User_Manual()
		os.Exit(1)
	}

	url := args[0]

	defer cleanup()

	if *versonflag {
		fmt.Println(version.Name, version.Version)
		return
	}

	if *selectorFlag != "" {
		results, err := scraper.ScrapeBySelector(url, *selectorFlag)
		if err != nil {
			fmt.Fprintln(writer, "error scraping by selector :", err)
			os.Exit(1)
		}

		if *jsonFlag {
			err := fileformatter.OutputAsJSON(writer, map[string]any{
				"selector": *selectorFlag,
				"count":    len(results),
				"results":  results,
			})
			if err != nil {
				fmt.Fprintln(writer, "error outputting as json :", err)
				os.Exit(1)
			}
			return
		}

		if *csvFlag {
			err := fileformatter.OutputAsCSV(writer, results)
			if err != nil {
				fmt.Fprintln(writer, "error outputting as csv :", err)
				os.Exit(1)
			}
			return
		}

		fmt.Fprintf(writer, "Result for selector \"%s\":\n\n", *selectorFlag)
		for i, r := range results {
			fmt.Fprintf(writer, "%d. %s\n", i+1, r)
		}

		fmt.Fprintf(writer, "\nTotal matches : %d\n", len(results))
		return
	}

	result, err := scraper.Scrape(url)
	if err != nil {
		fmt.Fprintln(writer, "error scraping the url :", err)
		os.Exit(1)
	}

	if titleFlag {
		fmt.Fprintln(writer, "Title : ", result.Title)
		return
	}

	if linksFlag {

		if *jsonFlag {
			err := fileformatter.OutputAsJSON(writer, map[string]any{
				"selector": url,
				"count":    len(result.Links),
				"results":  result.Links,
			})
			if err != nil {
				fmt.Fprintln(writer, "error outputting as json :", err)
				os.Exit(1)
			}
			return
		}

		if *csvFlag {
			err := fileformatter.OutputAsCSV(writer, result.Links)
			if err != nil {
				fmt.Fprintln(writer, "error outputting as csv :", err)
				os.Exit(1)
			}
			return
		}
		fmt.Fprintln(writer, "Links :")
		for i, link := range result.Links {
			fmt.Fprintf(writer, "%d. %s\n", i+1, link)
		}
		fmt.Fprintf(writer, "\nTotal links found : %d\n ", len(result.Links))
		return
	}

	fmt.Fprintln(writer, "title :", result.Title)
	fmt.Fprintln(writer, "links found : ", len(result.Links))

	fmt.Fprintln(writer, "\n review (first 5 links)")
	limit := 5
	if len(result.Links) < 5 {
		limit = len(result.Links)
	}

	for i := 0; i < limit; i++ {
		fmt.Fprintf(writer, "%d. %s\n", i+1, result.Links[i])
	}

}

func getWriter(outFile string) (io.Writer, func(), error) {
	if outFile == "" {
		return os.Stdout, func() {}, nil
	}

	f, err := os.Create(outFile)
	if err != nil {
		return nil, nil, err
	}

	return f, func() { f.Close() }, nil
}
