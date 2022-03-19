package mail

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/meiji163/gh-mail/pkg/encrypt"
)

func GetPublicKey(user, host string) (*rsa.PublicKey, error) {
	resp, err := http.Get(fmt.Sprintf("https://raw.%s/%s/inbox/main/public.pem", host, user))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("HTTP Error (%s): %d", resp.Request.URL, resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return encrypt.BytesToPublicKey(body)
}
