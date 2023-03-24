package main

import (
	"fmt"
	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
	"os"
)

func Run(rl *fn.ResourceList) (bool, error) {

	for index, item := range rl.Items {
		fmt.Printf("Item %d:\n---------------------------------------\n%s\n\n", index+1, item)
	}

	fmt.Printf("END OF LOOP\n")
	return true, nil
}

func main() {
	if err := fn.AsMain(fn.ResourceListProcessorFunc(Run)); err != nil {
		os.Exit(1)
	}
}
