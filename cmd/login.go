package cmd

import (
	"github.com/spf13/cobra"
	"wire_test/di"
	"wire_test/pkg/config"
)

var email = &flagDetails{
	name: "email",
	help: "Input your e-mail address",
}

var password = &flagDetails{
	name: "password",
	help: "Input your password",
}

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log into your gmail account",
	Long: `Lets you access your gmail account.
It has to be done before any of the other operations are.`,
	Run: func(cmd *cobra.Command, args []string) {
		authenticator, closeup, err := di.InitializeAuthenticator(config.GoogleSmtpAddress, config.StorageDirectory)
		defer closeup()
		if err != nil {
			cmd.PrintErrf("authenticator init failed: %v\n", err)
			return
		}
		promptForMissingFlags([]*flagDetails{email, password})
		err = authenticator.Login(*email.content.(*string), *password.content.(*string))
		if err != nil {
			cmd.PrintErrf("failed to authenticate: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	email.content = loginCmd.Flags().String(email.name, "", email.help)
	password.content = loginCmd.Flags().String(password.name, "", password.help)
}
