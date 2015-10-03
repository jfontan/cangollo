package main

import (
	"./cangallo"
	"fmt"
)

func main() {
	repo := cangallo.Repo{Path: "."}

	repo.Init()

	repo.AddImage("javi", cangallo.Image{SHA1: "javi"})

	fmt.Printf("%+v\n", repo.Index)

	repo.Marshal()
}
