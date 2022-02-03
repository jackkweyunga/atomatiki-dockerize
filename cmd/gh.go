/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"log"

	"github.com/spf13/cobra"

	"janjas/atomatiki/github"

	"janjas/atomatiki/config"
)

// ghCmd represents the gh command
var ghCmd = &cobra.Command{
	Use:   "gh",
	Short: "Gh gives you ability to use some specific gh commands",
	Long:  `...`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gh is called")
	},
}

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Pushes the local project to github.",
	Long:  "...",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gh push is called")
	},
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a remote project in github.",
	Long:  "...",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("gh create is called")

		data, err := config.Configuration()
		if err != nil {
			log.Fatal(err)
		}

		remoteName := data["config"].Template.RemoteName
		ghToken := data["config"].Docker.GhKey

		cBash := github.CreateRepoBash(remoteName, ghToken)
		c, err, f := config.BashCmdExec(cBash)
		if err != nil {
			log.Fatal(err)
		}

		output, err := c.CombinedOutput()
		if err != nil {
			log.Println(string(output))
			log.Fatal(err)
		}
		if f != nil {
			config.ClearTempFile(f)
		}

		log.Print(string(output))

	},
}

func init() {
	rootCmd.AddCommand(ghCmd)
	ghCmd.AddCommand(pushCmd, createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ghCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ghCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
