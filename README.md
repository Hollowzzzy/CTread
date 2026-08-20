# CTread

> A simple command-line ebook reader written in Go.

**CTread** is a terminal-based reading application designed to make managing and reading ebooks directly from the command line simple.

[![Go](https://img.shields.io/badge/Go-1.XX-00ADD8?logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

## Features

* Read ebooks directly from the terminal
* EPUB support
* Import books into a local registry
* Open registered books by name
* CLI interface
* Persistent configuration
* Local book registry
* Chapter-based reading

## Demo

<!-- TODO: Add a GIF showing the main CTread interface -->

![CTread demo](docs/gifs/main-demo.gif)

## Installation

### From source

Make sure you have Go installed, then clone the repository:

```bash
git clone https://github.com/Hollowzzzy/CTread.git
cd CTread
```

Build CTread:

```bash
go build -o CTread
```

You can then run it with:

```bash
./CTread
```

### Installing globally

You can install CTread somewhere on your `PATH`:

```bash
go install
```

Or copy the compiled binary to a directory such as:

```bash
cp CTread /opt/homebrew/bin/CTread
```

After that, you can run:

```bash
CTread
```

## Usage

### Import a book

```bash
CTread import "path/to/book.epub"
```

<!-- TODO: Add a GIF showing the import command -->

![Importing a book](docs/gifs/import.gif)

### Read a book

You can read a book by providing its path:

```bash
CTread read "path/to/book.epub"
```

If the book has been imported into the registry, you can also use its registered name:

```bash
CTread read "The Adventures of Sherlock Holmes"
```

<!-- TODO: Add a GIF showing reading a book -->

![Reading a book](docs/gifs/reading.gif)

### Read a specific chapter

```bash
CTread read "The Adventures of Sherlock Holmes" 5
```

## Supported Formats

Currently supported:

* EPUB

More formats are planned for future versions.

## Project Structure

A simplified overview of the project:

```text
ctread/
├── books/
│   ├── importers/
│   │   ├── epub.go
│   │   └── registry.go
│   └── reader/
│       └── reader.go
├── cmd/
│   ├── cli/
│   │   ├── config.go
│   │   ├── import.go
│   │   ├── read.go
│   │   ├── registry.go
│   │   └── root.go
│   └── func/
│       ├── config.go
│       ├── import.go
│       └── read.go
├── config/
│   └── config.go
├── .gitignore
├── go.mod
├── go.sum
├── main.go
└── README.md
```

The project is split into separate packages so that functionality such as importing, reading, configuration, and registry management can be developed independently.

## Configuration

CTread stores its configuration and book registry in the user's application configuration directory.

The registry contains information about imported books, including their names and file paths.

<!-- TODO: Add a GIF showing configuration/settings -->

![Configuration](docs/gifs/config.gif)

## Roadmap

* [x] Basic CLI
* [x] EPUB importing
* [x] Book registry
* [x] Chapter reading
* [ ] Improved reading experience
* [ ] TUI
* [ ] Reading progress
* [ ] Bookmarks
* [ ] Search within books
* [ ] More ebook formats
* [ ] Better library management
* [ ] Cross-platform packaging

## Contributing

Contributions, suggestions, and bug reports are welcome.

If you want to contribute:

```bash
git clone https://github.com/Hollowzzzy/CTread.git
cd CTread
go build
```

Make your changes, test them, and open a pull request.

## License

CTread is licensed under the MIT License.

See [LICENSE](LICENSE) for more information.

## Screenshots & GIFs

This section can be expanded as CTread develops.

<!-- TODO: Add a GIF demonstrating the full workflow -->

### Full workflow

![Full workflow](docs/gifs/full-workflow.gif)

<!-- TODO: Add a screenshot of the library -->

### Library

![Library](docs/images/library.png)

<!-- TODO: Add a screenshot of the reader -->

### Reader

![Reader](docs/images/reader.png)

---

Made with Go.
