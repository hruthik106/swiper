# **Swiper**

Swiper is a fast, lightweight **web data extraction CLI written in Go**.
It lets you quickly extract page titles, links, and custom HTML elements directly from the terminal..

## Features

- Extract page titles
- Extract all links from a page
- Scrape elements using CSS selectors
- Output in plain text, JSON, or CSV
- Save results to a file with --out
- Custom User-Agent support
- Clean, modular Go project structure

## Installation

### Install using go install (recommended)

```bash
go install github.com/hruthik106/swiper/cmd/swiper@latest
```

### Build from source

```bash
git clone https://github.com/YOUR_USERNAME/swiper.git
cd swiper
go build -o swiper ./cmd/swiper

```

### Run locally
```bash
./swiper --title https://example.com
```

## Usage

Important: Flags must come before the URL.

swiper [flags] <url>

Examples
#### Default summary
```bash
swiper https://example.com
```

#### Only title
```bash
swiper --title https://example.com
```

#### All links
```bash
swiper --links https://example.com
```

#### CSS selector scraping
```bash
swiper --selector "h1" https://example.com
```

#### JSON output
```bash
swiper --links --json https://example.com
```

#### CSV output
```bash
swiper --links --csv https://example.com
```

#### Write output to file
```bash
swiper --links --csv --out links.csv https://example.com
```

#### Custom User-Agent
```bash
swiper --title --ua "Mozilla/5.0" https://example.com
```

## Flags

| Flags | Description |
|------|------------|
| --title, -t | Show only the page title |
| --links, -l | Show all links |
| --selector  | Scrape by CSS selector |
| --json | Output as JSON |
| --csv | Output as CSV |
| --out | Write output to file |
| --ua|Override HTTP User-Agent|
| --version | Show version |

## Testing
#### Runs all tests 
```bash
go test ./...
```
#### Tests cover:
- Scraper logic
- Formatter (JSON / CSV)

## Contributing

Pull requests are welcome.

For major changes, please open an issue first.

## License

This project is licensed under the MIT License.


