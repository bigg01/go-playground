package main

import (
	"fmt"
	/*
		"log"
		"os"
		"os/exec"
		"path/filepath"
	*/
	"log"
	"os/exec"
)

type commando struct {
	cmdName string
	cmdArgs []string
}

func main() {

	cmd1 := new(commando)
	cmd1.cmdName = "ls"
	cmd1.cmdArgs = []string{"-la"}

	RunCMD(cmd1)
	//ExampleCmd_Output()

}

func ExampleCmd_Output() {
	out, err := exec.Command("ls").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The date is %s\n", out)
}

//Clone clones the specified repo in the executing directory
func RunCMD(c *commando) {
	cmd := exec.Command(c.cmdName, c.cmdArgs...)
	output, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", output)
	fmt.Printf("retcode: %s\n", err)


}
