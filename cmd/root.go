package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gen2brain/beeep"
	"github.com/spf13/cobra"
)

const defaultTitle = "Poptart"

var rootCmd = &cobra.Command{
	Use:   "poptart",
	Short: "Poptart is a simple CLI for generating notifications",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		alertTitle, _ := cmd.Flags().GetString("title")
		err := beeep.Alert(alertTitle, strings.Join(args, " "), "./cmd/poptart.png")
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
	rootCmd.PersistentFlags().StringP("title", "t", defaultTitle, "The title to set on the notification")
}

func initConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
}
