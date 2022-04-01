package Base

import (
	"fmt"
	"os"
	"os/exec"
)

func ExecCommand() {
	execCustom()
}

func execCustom() {
	env := os.Environ()
	procAttr := &os.ProcAttr{
		Env: env,
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}

	pid, err := os.StartProcess("/bin/ls", []string{"ls", "-la"}, procAttr)
	if err != nil {
		fmt.Printf("Errpr %v starting process!", err.Error())
		os.Exit(1)
	}

	fmt.Printf("process Pid:%v \n", pid)

	fmt.Println("=========================NewProcess=======================")

	cmd := exec.Command("open -a rdm")
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error %v executing command!", err)
		os.Exit(1)
	}

	fmt.Printf("The command is %v", cmd)

}
