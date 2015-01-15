package play_go

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func print_cmd(cmd *exec.Cmd) {
	fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}

func print_err(err error) {
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("==> BOOM: %s\n", err.Error()))
	}
}

func print_out(out []byte) {
	if len(out) > 0 {
		fmt.Printf("==> Output: %s\n", string(out))
	}
}

func run_with_return_code(cmd *exec.Cmd) *exec.ExitError {
	fmt.Printf("Command: %s\n", cmd.Args[0])
	output, result := cmd.CombinedOutput()
	if result != nil {
		fmt.Printf("Result: %s\n %s\n", result, output)
		fmt.Printf("Return code: %s\n", result.(*exec.ExitError))
	} else {
		fmt.Print("Succesful execution\n")
	}
	return result.(*exec.ExitError)
}

func IsValidDockerCommand(cmd *exec.Cmd) bool {
	if result := cmd.Run(); result != nil {
		switch result.(type) {
		default:
			//WTF?!?
			return false
		case *exec.ExitError:
			//The command failed because it's supported but missing arguments
			return true
		case *exec.Error:
			//Failed completely to execute!
			return false
		}
	} else {
		//The command succeeded when it shouldn't because it was not recognised
		return false
	}
}
