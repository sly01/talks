package functions

import (
	"fmt"
	"os/exec"
)

func CreateEC2() {

	cmd := exec.Command("terraform", "apply",  "-auto-approve", "/Users/aerkoc/terraform-workspace/assigment1")
	out, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print(string(out))
}

func DeleteEC2() {

        cmdDestroy := exec.Command("terraform", "destroy", "-auto-approve" , "/Users/aerkoc/terraform-workspace/assigment1")
        outDestroy, errDestroy := cmdDestroy.Output()

        if errDestroy != nil {
                fmt.Println(errDestroy.Error())
                return
        }

        fmt.Print(string(outDestroy))
}
