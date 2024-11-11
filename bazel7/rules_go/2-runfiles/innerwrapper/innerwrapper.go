package main

import (
	"fmt"
	"os/exec"

	"github.com/bazelbuild/rules_go/go/runfiles"
)

func main() {
	p, err := runfiles.Rlocation("bazel-example/go_version_/go_version")
	if err != nil {
		panic(err)
	}

	cmd := exec.Command(p)
	output, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Println("This is the inner wrapper")
	fmt.Println(string(output))
}
