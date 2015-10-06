package cangallo

import (
	//"fmt"
	"log"
	//"os"
	"os/exec"
	//"strconv"
	"encoding/json"
)

type QemuImg struct {
	Path string
}

type QemuImgInfo struct {
	VirtualSize         int64  `json:"virtual-size"`
	Filename            string `json:"filename"`
	ClusterSize         int    `json:"cluster-size"`
	Format              string `json:"format"`
	ActualSize          int64  `json:"actual-size"`
	FullBackingFilename string `json:"full-backing-filename"`
	BackingFilename     string `json:"backing-filename"`
	DirtyFlag           bool   `json:"dirty-flag"`
}

func (qemu_img *QemuImg) Execute(params []string) (text string, err error) {
	command_path := qemu_img.Path
	if command_path == "" {
		command_path = "/usr/bin/qemu-img"
	}

	cmd_text := []string{command_path}
	cmd_text = append(cmd_text, params...)

	log.Printf("Executing \"%v\"", cmd_text)

	cmd := exec.Command(command_path, params...)

	out, err := cmd.CombinedOutput()
	text = string(out)

	if err != nil {
		log.Printf("Error executing \"%v\"", cmd_text)
	}

	return text, err
}

func (qemu_img *QemuImg) Create(file string, size string) (err error) {
	params := []string{"create", "-f", "qcow2", file, size}

	text, err := qemu_img.Execute(params)

	if err != nil {
		log.Printf("Could not create the image. qemu-img message:")
		log.Printf(text)
	}

	return err
}

func (qemu_img *QemuImg) Info(file string) (info []QemuImgInfo, err error) {
	params := []string{"info", "--output", "json", "--backing-chain", file}

	text, err := qemu_img.Execute(params)

	if err != nil {
		log.Printf("Could get info from the image. qemu-img message:")
		log.Printf(text)

		return info, err
	}

	err = json.Unmarshal([]byte(text), &info)

	return info, err
}
