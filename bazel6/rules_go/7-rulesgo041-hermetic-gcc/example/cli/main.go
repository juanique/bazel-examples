package main

import (
	"time"

	"github.com/briandowns/spinner"
	"github.com/juanique/bazel-examples/example/capitalize"
)

func main() {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Suffix = capitalize.Capitalize("Hello there")
	s.Start()
	time.Sleep(4 * time.Second)
	s.Stop()
}
