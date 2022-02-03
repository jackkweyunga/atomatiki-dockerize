/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
)

var template = `
config:

  template:
    type: django
    remotename: "localname"
    localname: "remotename"
    port: 89

  docker:
    ghkey: "XXXXX"
    sshport: "22"
    sshhost: "example.com"
    sshusername: "user"
    dockerpassword: "password"
    dockerusername: "username"
    sshkey: "path/to/keyfile.txt"
`

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize an atomatiki.yaml configuration file",
	Long: `A atomatiki.yaml file will be created inn the path you are in.
	Edit the file specific to your needs.	
	Some initial data will be filled already. Some required fields will be commented.`,
	Run: func(cmd *cobra.Command, args []string) {

		// create a atomatiki.yaml file in the path provided.
		yfilePath := "atomatiki.yaml"

		err := ioutil.WriteFile(yfilePath, []byte(template), 0755)
		if err != nil {
			fmt.Println(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
