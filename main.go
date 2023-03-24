package main

import (
	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
	"github.com/matysiaq/kpt-examples/mutator"
	"os"
)

func main() {
	if err := fn.AsMain(fn.ResourceListProcessorFunc(mutator.Run)); err != nil {
		os.Exit(1)
	}
}
