package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gen2brain/beeep"
	"github.com/spf13/cobra"
)

const defaultTitle = "Poptart"
const defaultImage = "./cmd/poptart.png"

var rootCmd = &cobra.Command{
	Use:   "poptart",
	Short: "Poptart is a simple CLI for generating notifications",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		alertTitle, _ := cmd.Flags().GetString("title")
		alertImage, _ := cmd.Flags().GetString("image")

		err := beeep.Alert(alertTitle, strings.Join(args, " "), alertImage)

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
	rootCmd.PersistentFlags().StringP("image", "i", defaultImage, "Path to image")
}

func initConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
}
