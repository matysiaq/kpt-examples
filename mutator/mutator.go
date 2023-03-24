package mutator

import (
	"errors"
	"fmt"
	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
)

type MutatorCtx struct {
}

func (mc *MutatorCtx) GatherInfo(rl *fn.ResourceList) {

	// Test errors appending to the Results array
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

	// Access Object field using defined getters
	fmt.Printf("------------- Check KubeObject functions -------------\n")
	fmt.Printf("fn.KubeObject implements a lot of getters / setters / other functions, e.g.\n")
	for _, o := range rl.Items {
		fmt.Printf("\tGetAPIVersion: %s\n", o.GetAPIVersion())
		fmt.Printf("\tGetKind: %s\n", o.GetKind())
		fmt.Printf("\tIsLocalConfig: %t\n", o.IsLocalConfig())
		fmt.Printf("-------\n")
	}
	fmt.Printf("------------------------------------------------------\n")

	// Access Object nested fields using object.NesterString(fields ...string) method
	for _, o := range rl.Items {
		if o.GetAPIVersion() == "req.nephio.org/v1alpha1" && o.GetKind() == "Interface" {
			networkInstanceName, ok, err := o.NestedString("spec", "networkInstance", "name")
			if err != nil {
				rl.Results = append(rl.Results, fn.ErrorConfigObjectResult(err, o))
			}

			if ok {
				fmt.Printf("Accessed values of Interface [name=%v]: [%v, %v, %v]\n\n", o.GetName(), networkInstanceName, ok, err)
			}
		}
		if o.GetAPIVersion() == "req.nephio.org/v1alpha1" && o.GetKind() == "DataNetworkName" {
			pools, ok, err := o.NestedSlice("spec", "pools")
			if err != nil {
				rl.Results = append(rl.Results, fn.ErrorConfigObjectResult(err, o))
			}

			if ok {
				for _, pool := range pools {
					fmt.Printf("Accessed pool of DataNetworkName [name=%v]: [%v, %v, %v]\n\n", o.GetName(), pool.GetString("name"), ok, err)
				}
			}
		}
	}
}

func (mc *MutatorCtx) Mutation(rl *fn.ResourceList) {

}

func Run(rl *fn.ResourceList) (bool, error) {

	mc := &MutatorCtx{}

	mc.GatherInfo(rl)

	mc.Mutation(rl)

	return true, nil
}
