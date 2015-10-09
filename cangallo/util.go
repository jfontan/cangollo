package cangallo

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func OpenEditor(file_name string) (data []byte, err error) {
	cmd := exec.Command("/usr/bin/vim", file_name)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error editing file: %v\n", err)
		return nil, err
	}

	data, err = ioutil.ReadFile(file_name)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return nil, err
	}

	return data, err
}
