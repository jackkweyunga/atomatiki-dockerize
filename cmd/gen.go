/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"janjas/atomatiki/config"
)

// A function to create a project folder
func GenerateFolder(projectname string) error {
	if err := os.Mkdir(projectname, os.ModePerm); err != nil {
		return err
	}
	return nil
}

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "generates a local project folder",
	Long:  `The local project folder generated contains a DockerFile and a docker-compose.yaml file`,
	Run: func(cmd *cobra.Command, args []string) {

		data, err := config.Configuration()
		if err != nil {
			log.Fatal(err);
		}

		// STEPS

		// 1. CREATE PROJECT FOLDER
		if err = GenerateFolder(data["config"].Template.LocalName); err != nil {
			fmt.Print(err)
		} else {
			fmt.Print("Project created successfully")
		}

	},
}


func init() {
	rootCmd.AddCommand(genCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
