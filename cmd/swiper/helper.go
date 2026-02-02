package main

import (
	"flag"
	"fmt"
)

func Cli_User_Manual() {
	fmt.Println("Swiper – Web scraping CLI tool")
	fmt.Println("Usage:")
	fmt.Println("  swiper [flags] <url>")

	fmt.Println("Flags:")
	flag.PrintDefaults()

	fmt.Println("\nExamples:")
	fmt.Println("  swiper https://example.com")
	fmt.Println("  swiper --title https://example.com")
	fmt.Println("  swiper --links --csv --out links.csv https://example.com")
}
