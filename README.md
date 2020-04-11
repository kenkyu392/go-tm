# WIP: go-tm

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/kenkyu392/go-tm?status.svg)](https://godoc.org/github.com/kenkyu392/go-tm)
[![Go](https://img.shields.io/badge/go-1.14+-00ADD8.svg)](https://golang.org/)
[![Version](https://img.shields.io/badge/version-0.0.X-00A29C.svg)](README.md)

**:warning: This package is currently under development and features may change.**

L10N and T9N file format (TMX, TBX, XLIFF, etc.) library for [Go](https://golang.org/).

## Overview

go-tm is a package that provides functions and utilities for reading and writing translation and localization format XML files such as translation memories and term-bases.

The functions and utilities provided by this package are useful for developing programs that generate localization files such as translation memories and term-bases from existing translation files (XLSX, CSV, and other types of files).

- **Supports file formats such as TMX, TBX, and XLIFF.**
  - _Translation Memory eXchange_
    - LISA TMX 1.4
  - _[TermBase eXchange](https://www.tbxinfo.net/)_
    - LISA TBX 2.0 (ISO 30042:2008)
    - LISA TBX 3.0 (ISO 30042:2019)
  - _XML Localization Interchange File Format_
    - OASIS XLIFF 1.2
    - OASIS XLIFF 2.0
    - OASIS XLIFF 2.1
    - SDL XLIFF (OASIS XLIFF 1.2)
- **Over 300 [IETF BCP 47 language tags](docs/ietf-bcp-47-language-tags.md) are defined.**

Other useful functions and utilities will be added to the package.

## Installation

```
go get -u github.com/kenkyu392/go-tm
```

## Usage

```go
package main

func main() {

}
```

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## License

[MIT](LICENSE)
