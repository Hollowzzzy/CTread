# ctread

> A simple command-line ebook reader written in Go.

**ctread** is a terminal-based reading application designed to make managing and reading ebooks directly from the command line simple.

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

<!-- TODO: Add a GIF showing the main ctread interface -->

![ctread demo](docs/gifs/main-demo.gif)

## Installation

### From source

Make sure you have Go installed, then clone the repository:

```bash
git clone https://github.com/YOUR_USERNAME/ctread.git
cd ctread
```

Build ctread:

```bash
go build -o ctread
```

You can then run it with:

```bash
./ctread
```

### Installing globally

You can install ctread somewhere on your `PATH`:

```bash
go install
```

Or copy the compiled binary to a directory such as:

```bash
cp ctread /opt/homebrew/bin/ctread
```

After that, you can run:

```bash
ctread
```

## Usage

### Import a book

```bash
ctread import "path/to/book.epub"
```

<!-- TODO: Add a GIF showing the import command -->

![Importing a book](docs/gifs/import.gif)

### Read a book

You can read a book by providing its path:

```bash
ctread read "path/to/book.epub"
```

If the book has been imported into the registry, you can also use its registered name:

```bash
ctread read "Crime and Punishment"
```

<!-- TODO: Add a GIF showing reading a book -->

![Reading a book](docs/gifs/reading.gif)

### Read a specific chapter

```bash
ctread read "Crime and Punishment" 5
```

## Supported Formats

Currently supported:

* EPUB

More formats are planned for future versions.

## Project Structure

A simplified overview of the project:

```text
ctread/
├── cmd/
│   └── ...
├── config/
│   └── ...
├── importers/
│   └── ...
├── reader/
│   └── ...
├── registry/
│   └── ...
├── main.go
├── go.mod
└── README.md
```

The project is split into separate packages so that functionality such as importing, reading, configuration, and registry management can be developed independently.

## Configuration

ctread stores its configuration and book registry in the user's application configuration directory.

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
git clone https://github.com/YOUR_USERNAME/ctread.git
cd ctread
go build
```

Make your changes, test them, and open a pull request.

## License

ctread is licensed under the MIT License.

See [LICENSE](LICENSE) for more information.

## Screenshots & GIFs

This section can be expanded as ctread develops.

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
