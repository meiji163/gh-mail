package mail

import (
	"fmt"
	"path/filepath"

	"github.com/meiji163/gh-mail/cmd/util"
	"github.com/meiji163/gh-mail/pkg/encrypt"
	"github.com/spf13/cobra"
)

func NewCmdKeygen() *cobra.Command {
	return &cobra.Command{
		Use:   "keygen",
		Short: "generate encryption keys",
		RunE: func(cmd *cobra.Command, args []string) error {
			priv, pub := encrypt.GenerateKeys(4096)
			extPath := filepath.Join(util.ExtensionsDir(), "gh-mail")
			privPath := filepath.Join(extPath, "private.pem")
			pubPath := filepath.Join(extPath, "public.pem")

			if err := encrypt.WritePublicKeyPEM(pub, pubPath); err != nil {
				return err
			}

			if err := encrypt.WritePrivateKeyPEM(priv, privPath); err != nil {
				return err
			}

			fmt.Printf("%s\n%s", pubPath, privPath)
			return nil
		},
	}
}
