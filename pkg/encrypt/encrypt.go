package encrypt

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	pem "encoding/pem"
	"errors"
	"os"
)

var MaxMsg2048 int = 190

var InvalidPEMError error = errors.New("Invalid PEM")

func GenerateKeys(bits int) (*rsa.PrivateKey, *rsa.PublicKey) {
	priv, _ := rsa.GenerateKey(rand.Reader, bits)
	pub := &priv.PublicKey
	return priv, pub
}

func Encrypt(msg []byte, pub *rsa.PublicKey) ([]byte, error) {
	return rsa.EncryptOAEP(sha256.New(), rand.Reader, pub, msg, nil)
}

func Decrypt(cipher []byte, priv *rsa.PrivateKey) ([]byte, error) {
	return priv.Decrypt(nil, cipher, &rsa.OAEPOptions{Hash: crypto.SHA256})
}

// WritePublicKeyPEM writes RSA public key to pem file
func WritePublicKeyPEM(pub *rsa.PublicKey, fileName string) error {
	pubASN1, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0444)
	if err != nil {
		return err
	}
	defer file.Close()

	return pem.Encode(file,
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubASN1,
		})

	return err
}

// WritePrivateKeyPEM writes RSA private key to pem file
func WritePrivateKeyPEM(priv *rsa.PrivateKey, fileName string) error {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0400)
	if err != nil {
		return err
	}
	defer file.Close()

	return pem.Encode(file,
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(priv),
		})
}

// BytesToPrivateKey reads RSA private key from pem encoded bytes
func BytesToPrivateKey(priv []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(priv)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	var err error
	if enc {
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			return nil, err
		}
	}
	return x509.ParsePKCS1PrivateKey(b)
}

// BytesToPublicKey reads RSA private key from pem encoded bytes
func BytesToPublicKey(pemBytes []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(pemBytes)
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, InvalidPEMError
	}
	rsaKey, ok := key.(*rsa.PublicKey)
	if !ok {
		return nil, InvalidPEMError
	}
	return rsaKey, nil
}

func EncodeMsg(msg []byte, msgType string) []byte {
	return pem.EncodeToMemory(&pem.Block{
		Type:    "MESSAGE",
		Headers: map[string]string{"Encoding": msgType},
		Bytes:   msg,
	})
}

func DecodeMsg(msg []byte) *pem.Block {
	block, _ := pem.Decode(msg)
	return block
}
