# **Swiper**

Swiper is a fast and simple **web scraping CLI tool written in Go**.
It allows you to extract titles, links, and custom HTML elements
directly from the command line.

## Features

- Scrape page titles
- Extract all links
- Scrape elements using CSS selectors
- Output in plain text, JSON, or CSV
- Write output to a file using `--out`
- Clean, modular Go project structure

## Installation

### Install using go install (recommended)

```bash
go install github.com/hruthik106/swiper/cmd/swiper@v1.0.1
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

## Flags

| Flags | Description |
|------|------------|
| --title, -t | Show only the page title |
| --links, -l | Show all links |
| --selector  | Scrape by CSS selector |
| --json | Output as JSON |
| --csv | Output as CSV |
| --out | Write output to file |
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


