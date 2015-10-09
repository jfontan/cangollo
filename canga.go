package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/exec"
	//"strconv"
	"github.com/k0kubun/pp"

	"./cangallo"
)

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

	cmd := exec.Command("/usr/bin/vim", file_name)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error editing file: %v\n", err)
		os.Exit(-1)
	}

	text, err := ioutil.ReadFile(file_name)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(-1)
	}

	image := cangallo.Image{}

	err = yaml.Unmarshal(text, &image)
	if err != nil {
		fmt.Printf("Error parsing yaml: %v\n", err)
		os.Exit(-1)
	}

	pp.Print(image)

	repo.AddImage("test", image)

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
