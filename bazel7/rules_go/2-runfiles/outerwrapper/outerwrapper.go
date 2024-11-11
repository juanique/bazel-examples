package main

import (
	"fmt"
	"os/exec"

	"github.com/bazelbuild/rules_go/go/runfiles"
)

func main() {
	p, err := runfiles.Rlocation("bazel-example/innerwrapper/innerwrapper_/innerwrapper")
	if err != nil {
		panic(err)
	}

	cmd := exec.Command(p)
	output, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Println("This is the outer wrapper")
	fmt.Println(string(output))
}
