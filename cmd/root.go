package main

import (
	"fmt"
	"os"

	"github.com/gen2brain/beeep"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "poptart",
	Short: "Poptart is a simple CLI for generating notifications",
	Long:  ``,
	// Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		err := beeep.Alert("PopTart", args[0], "./cmd/poptart.png")
		cobra.CheckErr(err)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "Author name for copyright attribution")
}

func initConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
}
