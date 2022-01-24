/*
Copyright © 2022 Ross Bryan

*/
package main

import (
	"log"

	"github.com/bryanro92/dprofile/cmd"
	version "github.com/bryanro92/dprofile/pkg/version"
)

func main() {
	log.Print("running version:", version.GitCommit)
	cmd.Execute()
}
