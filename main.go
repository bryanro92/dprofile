/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
