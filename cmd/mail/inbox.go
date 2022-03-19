package mail

import (
	"fmt"
	"os"
	"strings"

	"github.com/meiji163/gh-mail/cmd/util"
	"github.com/meiji163/gh-mail/pkg/encrypt"
	"github.com/meiji163/gh-mail/pkg/issues"
	"github.com/spf13/cobra"
)

func NewCmdInbox(cfg util.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "inbox",
		Short: "read messages in your inbox",
		RunE: func(cmd *cobra.Command, args []string) error {
			return inboxRun(cfg)
		},
	}
	return cmd
}

func inboxRun(cfg util.Config) error {
	issues, err := issues.GetIssues(cfg.Login(), "inbox")
	if err != nil {
		return err
	}

	b, err := os.ReadFile(util.PrivateKeyPath())
	if err != nil {
		return err
	}
	privKey, err := encrypt.BytesToPrivateKey(b)
	if err != nil {
		return err
	}

	for i, issue := range issues {
		if !strings.HasPrefix(issue.Title, "gh-mail:") {
			continue
		}

		title := strings.TrimPrefix(issue.Title, "gh-mail: ")
		fmt.Printf("%d from: %s ------ %s\n", i, issue.User.Login, title)

		block := encrypt.DecodeMsg([]byte(issue.Body))
		msg, err := encrypt.Decrypt(block.Bytes, privKey)
		if err != nil {
			fmt.Print("\nERROR Decrypting message\n\n")
		} else {
			fmt.Printf("\n%s\n\n", string(msg))
		}
	}
	return nil
}
