/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"wire_test/di"
	"wire_test/pkg/config"
)

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		authenticator, closeup, err := di.InitializeAuthenticator(config.GoogleSmtpAddress, config.StorageDirectory)
		defer closeup()
		if err != nil {
			cmd.Println(err)
			return
		}
		err = authenticator.Logout()
		if err != nil {
			cmd.Println(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}
