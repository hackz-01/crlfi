package main

import "github.com/hackz-01/crlfi/internal/runner"

func main() {
	options = runner.ParseOptions()
	runner.New(options)
}
