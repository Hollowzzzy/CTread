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

## Installation

### From source

Make sure you have Go installed, then clone the repository:

```bash
git clone https://github.com/Hollowzzzy/ctread.git
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

You can install ctread directly using Go:

```bash
go install github.com/Hollowzzzy/CTread@latest
```

Alternatively, you can copy a compiled binary to a directory on your `PATH`:

```bash
cp ctread /opt/homebrew/bin/ctread
```

After that, you can run:

```bash
ctread
```

## Building for multiple platforms

ctread includes a `build.sh` script for building binaries for multiple operating systems and architectures.

Make sure the script is executable:

```bash
chmod +x build.sh
```

Then run:

```bash
./build.sh
```

The script builds ctread for several common platforms and architectures and places the resulting binaries in the `dist/` directory.

Currently supported builds include:

* macOS ARM64
* macOS AMD64
* Linux ARM64
* Linux AMD64
* Windows ARM64
* Windows AMD64
* FreeBSD ARM64
* FreeBSD AMD64

The resulting files will look similar to:

```text
dist/
├── ctread-darwin-arm64
├── ctread-darwin-amd64
├── ctread-linux-arm64
├── ctread-linux-amd64
├── ctread-windows-arm64.exe
├── ctread-windows-amd64.exe
├── ctread-freebsd-arm64
└── ctread-freebsd-amd64
```

## Usage

### Import a book

```bash
ctread import "path/to/book.epub"
```

![Importing a book](docs/examples/gifs/ImportDemo.gif)

### Registry

ctread includes a local book registry that allows imported books to be opened by their registered name instead of specifying their full file path.

You can view the current registry with:

```bash
ctread registry
```

![Book registry](docs/examples/gifs/RegistryDemo.gif)

Once a book has been imported, you can open it using its registered name:

```bash
ctread read "The Adventures of Sherlock Holmes"
```

The registry stores the book's name and file path in ctread's local application configuration directory.

### Read a book

You can read a book by providing its path:

```bash
ctread read "path/to/book.epub"
```

If the book has been imported into the registry, you can also use its registered name:

```bash
ctread read "The Adventures of Sherlock Holmes"
```

![Reading a book](docs/examples/gifs/ReadDemo.gif)

### Read a specific chapter

```bash
ctread read "The Adventures of Sherlock Holmes" 5
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
├── docs/examples/
│   ├── gifs/
│   │   ├── ImportDemo.gif
│   │   ├── ReadDemo.gif
│   │   └── RegistryDemo.gif
│   └── tapes/
│       ├── import.tape
│       ├── read.tape
│       └── registry.tape
├── libpgloss/
│   └── styles.go
├── .gitignore
├── build.sh
├── go.mod
├── go.sum
├── LICENSE
├── main.go
└── README.md
```

## Configuration

ctread stores its configuration and book registry in the user's application configuration directory.

The registry contains information about imported books, including their names and file paths.

Currently, there are no configuration settings that actually do anything, but the book registry does work.

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
git clone https://github.com/Hollowzzzy/ctread.git
cd ctread
```

Build the project:

```bash
go build
```

Or build binaries for multiple platforms:

```bash
chmod +x build.sh
./build.sh
```

Make your changes, test them, and open a pull request.

## License

ctread is licensed under the MIT License.

See [LICENSE](LICENSE) for more information.

---

Made with Go.
