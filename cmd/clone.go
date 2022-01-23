/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/bryanro92/dprofile/pkg/clone"
	"github.com/spf13/cobra"
)

// cloneCmd represents the clone command
var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clones repositories to a directory.",
	Long: `Able to take a list of repositories to clone
or clone all repositories of a given user.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		err := clone.Clone(ctx, args)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		fmt.Println("clone finished")
	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cloneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cloneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
