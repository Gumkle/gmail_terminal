/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
	"wire_test/di"
	"wire_test/pkg/config"
)

var to = &flagDetails{
	name: "to",
	help: "email address of the recipent",
}

var subject = &flagDetails{
	name: "subject",
	help: "subject of the email",
}

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		sender, closeup, err := di.InitializeSender(config.GoogleSmtpAddress, config.StorageDirectory)
		defer closeup()
		if err != nil {
			cmd.Println(err)
			return
		}
		promptForMissingFlags([]*flagDetails{to, subject})
		content, err := promptForEmailContent()
		if err != nil {
			cmd.Println(err)
			return
		}
		err = sender.Send(*to.content.(*string), *subject.content.(*string), content)
		if err != nil {
			cmd.Println(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)

	to.content = sendCmd.Flags().String(to.name, "", to.help)
	subject.content = sendCmd.Flags().String(subject.name, "", subject.help)
}

func promptForEmailContent() ([]byte, error) {
	fmt.Println("Please enter the email content. When done, send EOF (ctrl+D):")
	all, err := io.ReadAll(bufio.NewReader(os.Stdin))
	if err != nil {
		return nil, fmt.Errorf("error reading the contents: %w", err)
	}
	return all, nil
}
