package main

import (
	"errors"
	"fmt"
	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
	"os"
)

func Run(rl *fn.ResourceList) (bool, error) {

	for i, item := range rl.Items {
		fmt.Printf("Item %d:\n---------------------------------------\n%s\n\n", i+1, item)
		if i == 2 {
			rl.Results = append(rl.Results, fn.ErrorResult(errors.New("index 2 not accepted")))
		}
		if i == 3 {
			rl.Results = append(rl.Results, fn.ErrorConfigObjectResult(errors.New("not allowed"), rl.Items[i]))
		}
	}
	fmt.Printf("END OF LOOP\n")

	fmt.Printf("------------- Check KubeObject functions -------------\n")
	fmt.Printf("fn.KubeObject implements a lot of getters / setters / other functions, e.g.\n")
	for _, o := range rl.Items {
		fmt.Printf("\tGetAPIVersion: %s\n", o.GetAPIVersion())
		fmt.Printf("\tGetKind: %s\n", o.GetKind())
		fmt.Printf("\tIsLocalConfig: %t\n", o.IsLocalConfig())
		fmt.Printf("-------\n")
	}
	fmt.Printf("------------------------------------------------------\n")

	return true, nil
}

func main() {
	if err := fn.AsMain(fn.ResourceListProcessorFunc(Run)); err != nil {
		os.Exit(1)
	}
}
