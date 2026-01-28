# Swiper 🧹

Swiper is a simple and fast web scraping CLI tool written in Go.

## Features (v1)
- Fetch webpage title
- Extract all links
- Smart default summary
- User-Agent support
- Relative → absolute URL handling

## Installation

Build from source:

```bash
go build -o swiper ./cmd/swiper


usage 

swiper [flag] <url>

example 

swiper https://example.com
swiper --title https://example.com
swiper --links https://example.com



