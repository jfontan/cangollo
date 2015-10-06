package main

import (
	//"fmt"
	"github.com/codegangsta/cli"
	"os"
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

	//fmt.Printf("%+v\n", info)
	pp.Print(info)
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
}

func main() {
	app := cli.NewApp()
	app.Name = "canga"
	app.Commands = Commands
	app.Run(os.Args)
}
