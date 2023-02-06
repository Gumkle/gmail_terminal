package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"wire_test/pkg/auth"
	"wire_test/pkg/config"
)

type flagDetails struct {
	content interface{}
	name    string
	help    string
}

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
		promptForMissingData([]*flagDetails{email, password})
		authenticator, err := auth.InitializeAuthenticator(config.GoogleSmtpAddress, config.StorageDirectory)
		if err != nil {
			cmd.PrintErrf("authenticator init failed: %v\n", err)
			return
		}
		err = authenticator.Login(*email.content.(*string), *password.content.(*string))
		if err != nil {
			cmd.PrintErrf("failed to authenticate: %v\n", err)
		}
	},
}

func promptForMissingData(fd []*flagDetails) {
	for _, details := range fd {
		if *details.content.(*string) != "" {
			continue
		}
		for {
			fmt.Printf("%s: ", details.help)
			_, err := fmt.Scanf("%s", details.content)
			if err == nil {
				break
			}
		}
	}
}

func init() {
	rootCmd.AddCommand(loginCmd)
	email.content = loginCmd.Flags().String(email.name, "", email.help)
	password.content = loginCmd.Flags().String(password.name, "", password.help)
}
