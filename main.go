package main

import (
	"fmt"
	"os"

	"github.com/meiji163/gh-mail/cmd/mail"
	"github.com/meiji163/gh-mail/cmd/util"
	"github.com/spf13/cobra"
)

func rootCmd(cfg util.Config) *cobra.Command {
	mailCmd := &cobra.Command{
		Use:   "mail",
		Short: "send encrypted messages",
	}
	mailCmd.AddCommand(mail.NewCmdSend())
	mailCmd.AddCommand(mail.NewCmdKeygen())
	mailCmd.AddCommand(mail.NewCmdInbox(cfg))
	return mailCmd
}

func main() {
	config := util.NewConfig()

	cmd := rootCmd(config)
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
