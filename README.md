# pix2ansi

Convert images to ANSI escape code art for your terminal.

Inspired by [px2ansi](https://github.com/Nellousan/px2ansi) by Nellousan — rewritten in Go for easy cross-platform distribution as a single binary.

## How it works

`pix2ansi` reads an image and maps every two rows of pixels into a single terminal row using the `▀` and `▄` Unicode block characters, combining foreground and background ANSI 24-bit colors to effectively double the vertical resolution. Transparent pixels are rendered as blank space.

## Installation

**Download a pre-built binary:**

Head to the [Releases](https://github.com/MawCeron/pix2ansi/releases) page and download the binary for your platform. No dependencies required.

**From source:**
```bash
git clone https://github.com/MawCeron/pix2ansi
cd pix2ansi
go build -o pix2ansi .
```

**Using `go install`:**
```bash
go install github.com/MawCeron/pix2ansi@latest
```

## Usage

```
pix2ansi [options] <filename>
```

**Options:**

| Flag | Description |
|------|-------------|
| `-o`, `--output <file>` | Write output to a file instead of stdout |

**Examples:**

Print an image directly to the terminal:
```bash
pix2ansi image.png
```

Save the ANSI output to a file:
```bash
pix2ansi image.png -o output.ans
```

Pipe into less with color support:
```bash
pix2ansi image.png | less -R
```

## Supported formats

- PNG
- JPEG
- WebP

All images are internally converted to RGBA before processing, so transparency is fully supported.

## Requirements

- Go 1.21+
- A terminal with [24-bit true color](https://github.com/termstandard/colors) support

## License

MIT