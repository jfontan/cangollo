
package main

import (
    "fmt"
    "cangallo"
)

func main() {
    repo := Repo{Path: "."}
    repo.Init()

    repo.AddImage("javi", Image{SHA1: "javi"})


    fmt.Printf("%+v\n", repo.Index)

    repo.Marshal()
}


