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
    - [x] LISA TMX 1.4
  - _[TermBase eXchange](https://www.tbxinfo.net/)_
    - [ ] LISA TBX 2.0 (ISO 30042:2008)
    - [ ] LISA TBX 3.0 (ISO 30042:2019)
  - _XML Localization Interchange File Format_
    - [x] OASIS XLIFF 1.2
    - [ ] OASIS XLIFF 2.0
    - [ ] OASIS XLIFF 2.1
    - [ ] SDL XLIFF (OASIS XLIFF 1.2)
- **Over 300 [IETF BCP 47 language tags](docs/ietf-bcp-47-language-tags.md) are defined.**

Other useful functions and utilities will be added to the package.

## Installation

```
go get -u github.com/kenkyu392/go-tm
```

## Usage

_XLIFF 1.2 :_

```go
package main

import (
	"github.com/kenkyu392/go-tm"
	xliff "github.com/kenkyu392/go-tm/xliff/1.2"
)

func main() {
	xlf, err := xliff.New(
		// Set the source and target languages for the XLIFF file.
		xliff.SourceLanguageOption(tm.Tag_jaJP),
		xliff.TargetLanguageOption(tm.Tag_enUS),
		// You can specify the encoding of the XML file.
		xliff.UseUTF8XMLEncodingOption(),
	)
	if err != nil {
		panic(err)
	}

	// Add a translation unit to the file.
	xlf.AddTransUnit(
		// Translation Unit ID.
		"ID001",
		// Source text and comments.
		&xliff.Unit{
			Text: "吾輩は猫である",
			Notes: []string{
				"著者：夏目漱石",
				"発行日：1905年～1906年",
			},
		},
		// Target text and comments.
		&xliff.Unit{
			Text: "I Am a Cat",
			Notes: []string{
				"Author‎: ‎Natsume Sōseki",
				"Publication date: 1905–1906",
			},
		},
	)
	raw, err := tm.Encode(xlf)
	if err != nil {
		panic(err)
	}
	println(string(raw))
  /*
	<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
	<xliff xmlns="urn:oasis:names:tc:xliff:document:1.2" xml:space="preserve" version="1.2">
	  <file date="2020-01-01T00:00:00Z" original="original.xlf" datatype="plaintext" source-language="ja-JP" target-language="en-US">
	    <header>
	      <tool tool-id="go-tm" tool-name="GoTM XLIFF" tool-version="0.1.0"></tool>
	    </header>
	    <body>
	      <trans-unit xml:space="preserve" id="ID001">
	        <source xml:lang="ja-JP">吾輩は猫である</source>
	        <target xml:lang="en-US">I Am a Cat</target>
	        <note xml:lang="ja-JP">著者：夏目漱石</note>
	        <note xml:lang="ja-JP">発行日：1905年～1906年</note>
	        <note xml:lang="en-US">Author‎: ‎Natsume Sōseki</note>
	        <note xml:lang="en-US">Publication date: 1905–1906</note>
	      </trans-unit>
	    </body>
	  </file>
	</xliff>
	*/
}
```

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## License

[MIT](LICENSE)
