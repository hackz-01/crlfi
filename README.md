# CRLFI

A fast tool to scan CRLF vulnerability written in Go

<img src="https://user-images.githubusercontent.com/25837540/90128972-fc3bdf00-dd91-11ea-8c3b-0d6f4e8c6ba3.png" height="350">

---

## Resources

- [Installation](#installation)
	- [from Binary](#from-binary)
	- [from Source](#from-source)
	- [from GitHub](#from-github)
- [Usage](#usage)
	- [Basic Usage](#basic-usage)
	- [Flags](#flags)
	- [Target](#target)
		- [Single URL](#single-url)
		- [URLs from list](#urls-from-list)
		- [from Stdin](#from-stdin)
	- [Method](#method)
	- [Data](#data)
	- [Adding Headers](#adding-headers)
	- [Using Proxy](#using-proxy)
	- [Concurrency](#concurrency)
	- [Silent](#silent)
	- [Verbose](#verbose)
	- [Version](#version)
	- [Library](#library)

## Installation

### from Binary

The installation is easy. You can download a prebuilt binary from [releases page](https://github.com/hackz-01/crlfi/releases), unpack and run! or with

```bash
curl -sSfL https://git.io/crlfi | sh -s -- -b /usr/local/bin
```

### from Source

If you have go1.13+ compiler installed and configured:

```bash
GO111MODULE=on go install github.com/hackz-01/crlfi/cmd/crlfi@latest
```

In order to update the tool, you can use `-u` flag with go get command.

### from GitHub

```bash
git clone https://github.com/hackz-01/crlfi
cd crlfi/cmd/crlfi
go build .
mv crlfi /usr/local/bin
```

## Usage

### Basic Usage

Simply, crlfi can be run with:

```bash
crlfi -u "http://target"
```

### Flags

```bash
crlfi -h
```

This will display help for the tool. Here are all the switches it supports.

| Flag             	| Description                                    	|
|------------------	|------------------------------------------------	|
| -u, --url        	| Define single URL to fuzz                      	|
| -l, --list       	| Fuzz URLs within file                          	|
| -X, --method     	| Specify request method to use _(default: GET)_ 	|
| -o, --output     	| File to save results                              |
| -d, --data       	| Define request data                            	|
| -H, --header     	| Pass custom header to target                   	|
| -x, --proxy      	| Use specified proxy to fuzz                       |
| -c, --concurrent 	| Set the concurrency level _(default: 25)_      	|
| -s, --silent     	| Silent mode                                    	|
| -v, --verbose    	| Verbose mode                                   	|
| -V, --version    	| Show current crlfi version                   	|
| -h, --help       	| Display its help                               	|

### Target

You can define a target in 3 ways:

#### Single URL

```bash
crlfi -u "http://target"
```

#### URLs from list

```bash
crlfi -l /path/to/urls.txt
```

#### from Stdin

In case you want to chained with other tools.

```bash
subfinder -d target -silent | httpx -silent | crlfi
```

### Method

By default, crlfi makes requests with `GET` method.
If you want to change it, you can use the `-X` flag.

```bash
crlfi -u "http://target" -X "GET"
```

### Output

You can also save fuzzing results to a file with `-o` flag.

```bash
crlfi -l /path/to/urls.txt -o /path/to/results.txt
```

### Data

If you want to send a data request using POST, DELETE. PATCH or other methods, you just need to use `-d` flag.

```bash
crlfi -u "http://target" -X "POST" -d "data=body"
```

### Adding Headers

May you want to use custom headers to add cookies or other header parts.

```bash
crlfi -u "http://target" -H "Cookie: ..." -H "User-Agent: ..."
```

### Using Proxy

Using a proxy, proxy string can be specified with a `protocol://` prefix to specify alternative proxy protocols.

```bash
crlfi -u "http://target" -x http://127.0.0.1:8080
```

### Concurrency

Concurrency is the number of fuzzing at the same time. Default value CRLFI provide is `25`, you can change it by using `-c` flag.

```bash
crlfi -l /path/to/urls.txt -c 50
```

### Silent

If you activate this silent mode with the `-s` flag, you will only see vulnerable targets.

```bash
crlfi -l /path/to/urls.txt -s | tee vuln-urls.txt
```

### Verbose

Unlike silent mode, it will display error details if there is an error with the `-v` flag.

```bash
crlfi -l /path/to/urls.txt -v
```

### Version

To display the current version of CRLFI with the `-V` flag.

```bash
crlfi -V
```

### Library

You can use CRLFI as a library.

```go
package main

import (
	"fmt"

	"github.com/dwisiswant0/crlfi/pkg/crlfi"
)

func main() {
	target := "http://target"
	method := "GET"

	// Generates a potentially CRLF vulnerable URLs
	for _, url := range crlfi.GenerateURL(target) {
		// Scan against target
		vuln, err := crlfi.Scan(url, method, "", []string{}, "")
		if err != nil {
			panic(err)
		}

		if vuln {
			fmt.Printf("VULN! %s\n", url)
		}
	}
}
