package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/codegangsta/cli"
	"github.com/k0kubun/pp"
	"gopkg.in/yaml.v2"

	"./cangallo"
)

func GenericError(err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(-1)
	}
}

func CreateCommand(c *cli.Context) {
	path := c.Args()[0]
	size := c.Args()[1]

	qemuImg := cangallo.QemuImg{}
	qemuImg.Create(path, size)
}

func InfoCommand(c *cli.Context) {
	path := c.Args()[0]

	qemuImg := cangallo.QemuImg{}
	info, _ := qemuImg.Info(path)

	pp.Print(info)
}

func AddCommand(c *cli.Context) {
	image_file := c.Args()[0]

	repo := cangallo.Repo{}
	repo.Init()

	file, err := ioutil.TempFile("/tmp", "canga-")
	if err != nil {
		fmt.Printf("Can not create tempfile: %v\n", err)
		os.Exit(-1)
	}

	file.Write([]byte(cangallo.BasicImageText))

	file_name := file.Name()
	file.Close()

	data, err := cangallo.OpenEditor(file_name)
	GenericError(err)

	image := cangallo.Image{}

	err = yaml.Unmarshal(data, &image)
	if err != nil {
		fmt.Printf("Error parsing yaml: %v\n", err)
		os.Exit(-1)
	}

	temp_image_file, err := ioutil.TempFile(".", "canga-")
	if err != nil {
		fmt.Printf("Can not create tempfile: %v\n", err)
		os.Exit(-1)
	}

	temp_file_name := temp_image_file.Name()
	temp_image_file.Close()

	qemuImg := cangallo.QemuImg{}
	qemuImg.Clone(image_file, temp_file_name)

	sha1, err := cangallo.CalculateSHA1(temp_file_name)
	if err != nil {
		fmt.Printf("Error calculating SHA1: %v\n", err)
		os.Exit(-1)
	}

	dest_file_name := fmt.Sprintf("%s.qcow2", sha1)
	os.Rename(temp_file_name, dest_file_name)

	info, err := qemuImg.Info(dest_file_name)

	image.SHA1 = sha1
	image.TotalSize = info[0].ActualSize
	image.Size = info[0].VirtualSize
	image.Time = time.Now()

	pp.Print(image)

	repo.AddImage(sha1, image)

	pp.Print(repo.Index)

	repo.SaveIndex()
}

func ListCommand(c *cli.Context) {
	repo := cangallo.Repo{}
	repo.Init()

	pp.Print(repo)
}

var Commands = []cli.Command{
	{
		Name:  "test",
		Usage: "this is a test",
		Action: func(c *cli.Context) {
			println("test!")

			if c.String("wow") != "" {
				println(c.String("wow"))
			}
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "wow",
				Usage: "much wow",
			},
		},
	},
	{
		Name:   "create",
		Usage:  "create a new image",
		Action: CreateCommand,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "parent,p",
				Usage: "path of the parent image",
			},
		},
	},
	{
		Name:   "info",
		Usage:  "get info from an image",
		Action: InfoCommand,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "parent,p",
				Usage: "path of the parent image",
			},
		},
	},
	{
		Name:   "add",
		Usage:  "add a new image to the repository",
		Action: AddCommand,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "parent,p",
				Usage: "path of the parent image",
			},
		},
	},
	{
		Name:   "list",
		Usage:  "list images in the repository",
		Action: ListCommand,
	},
}

func main() {
	app := cli.NewApp()
	app.Name = "canga"
	app.Commands = Commands
	app.Run(os.Args)
}
