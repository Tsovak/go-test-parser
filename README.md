#  Beautify the Golang test output

[![License](https://img.shields.io/badge/license-GPLv3-blue.svg)](http://www.gnu.org/licenses/gpl-3.0.html)

License: GPL v3. [GPL License](http://www.gnu.org/licenses)

![tests](https://github.com/Tsovak/go-test-parser/workflows/tests/badge.svg)

## Introduction

This is the CLI fot make go test output more readable. 
The CLI aggregate all mixed test output and displays the tests with their output.

### Requirements

    1. Go 1.13 
    2. Make
    
## Build
Run the command in the project base dir 

    make build 
    
## Usage 

```bash
go-test-parser is a simple CLI tool for generating the UI report from go test output

Usage:
  ./go-test-parser <go test -json -v output> [flags]

Examples:
./go-test-parser ./test.log -v
 go test -json -v ./... | go-test-parser -o ./report_directory

Flags:
  -h, --help               Help for application
      --http int32         HTTP address to serve (default 8080)
  -s, --includeSucceeded   Include the successful test results or not
  -o, --output string      The report output directory (default "./report")
  -v, --verbose            Print verbose output
  -w, --web                Start only web serer for report displaying
```

1. Generate report to directory: `go-test-parser -o ./report_directory test.log`
2. Generate report and display in the browser via url http://localhost/report: `go-test-parser  -w --http 80 test.log`

###### Demo
![Report Demo](demo/demo.gif)

