package main

import "github.com/hackz-01/crli/internal/runner"

func main() {
	options := runner.ParseOptions()
	runner.New(options)
}
