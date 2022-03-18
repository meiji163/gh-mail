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

var MaxMsgBytes int = 190

var InvalidPEMError error = errors.New("Invalid PEM")

func GenerateKeys() (*rsa.PrivateKey, *rsa.PublicKey) {
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}
	pub := &priv.PublicKey
	return priv, pub
}

func Encrypt(msg []byte, pub *rsa.PublicKey) ([]byte, error) {
	return rsa.EncryptOAEP(sha256.New(), rand.Reader, pub, msg, nil)
}

func Decrypt(cipher []byte, priv *rsa.PrivateKey) ([]byte, error) {
	return privateKey.Decrypt(nil, cipher, &rsa.OAEPOptions{Hash: crypto.SHA256})
}

func WritePublicKeyPEM(pub *rsa.PublicKey, fileName string) error {
	pubASN1, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	return pem.Encode(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubASN1,
	}, file)

	return err
}

func WritePrivateKeyPEM(priv *rsa.PrivateKey, fileName string) error {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	return pem.Encode(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(priv),
	}, file)
}

func ReadPublicKeyPEM(fileName string) (*rsa.PublicKey, error) {
	pemBytes, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	block, err := pem.Decode(pemBytes)
	if err != nil {
		return nil, err
	}

	if !x509.IsEncryptedPEMBlock(block) {
		return nil, InvalidPEMError
	}

	decryptedBlock, err = x509.DecryptPEMBlock(block, nil)
	if err != nil {
		return nil, InvalidPEMError
	}

	return x509.ParsePKIXPublicKey(decryptedBlock)
}
