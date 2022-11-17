package util

import (
	"log"
	"os/exec"
)

func ExecCommand(cmd string, args ...string) (string, error) {
	c := exec.Command(cmd, args...)
	output, err := c.CombinedOutput()
	log.Printf("exec command %v: output: %v, err: %v", cmd, string(output), err)
	return string(output), err
}

func ExecShell(cmd string) (string, error) {
	output, err := ExecCommand("sh", "-c", cmd)
	log.Printf("exec shell %v: output: %v, err: %v", cmd, output, err)
	return output, err
}
