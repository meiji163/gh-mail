package mail

import (
	"errors"
	"fmt"

	"github.com/meiji163/gh-mail/pkg/encrypt"
	"github.com/meiji163/gh-mail/pkg/issues"
	"github.com/spf13/cobra"
)

type SendOptions struct {
	Recipient string
	Host      string
	Title     string
	Body      string
	doEncrypt bool
}

func NewCmdSend() *cobra.Command {
	opts := &SendOptions{
		Host:      "github.com",
		doEncrypt: true,
	}
	cmd := &cobra.Command{
		Use:   "send",
		Short: "send a message",
		RunE: func(cmd *cobra.Command, args []string) error {
			return sendRun(opts)
		},
	}
	cmd.Flags().StringVarP(&opts.Recipient, "recipient", "r", "", "user to send to")
	cmd.Flags().StringVarP(&opts.Title, "title", "t", "", "title of message")
	cmd.Flags().StringVarP(&opts.Body, "body", "b", "", "body of message")
	return cmd
}

func sendRun(opt *SendOptions) error {
	pub, err := GetPublicKey(opt.Recipient, opt.Host)
	if err != nil {
		return err
	}

	cipher, err := encrypt.Encrypt([]byte(opt.Body), pub)
	if err != nil {
		return err
	}

	blockMsg := encrypt.EncodeMsg(cipher, "base64")
	if blockMsg == nil {
		return errors.New("error encoding message")
	}

	return issues.CreateIssue(
		&issues.Issue{
			Title: fmt.Sprintf("gh-mail: %s", opt.Title),
			Body:  string(blockMsg),
		}, opt.Recipient, "inbox")
}
