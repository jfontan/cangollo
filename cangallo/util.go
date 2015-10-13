package cangallo

import (
	"crypto/sha1"
	"fmt"
	"io"
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

func CalculateSHA1(file_name string) (sha1_text string, err error) {
	file, err := os.Open(file_name)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return "", err
	}

	buffer := make([]byte, 1024)
	hash := sha1.New()

	var count int

	for ; err == nil; count, err = file.Read(buffer) {
		if count == 0 {
			continue
		}

		hash.Write(buffer[:count])
	}

	if err != nil && err != io.EOF {
		fmt.Printf("Error opening file: %v\n", err)
		return "", err
	}

	// Last error value is io.EOF. Clear it before returning.
	err = nil

	sha1_binary := hash.Sum(nil)
	sha1_text = fmt.Sprintf("%x", sha1_binary)

	return sha1_text, err
}
