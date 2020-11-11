package transactions

import (
	"crypto/ed25519"
	"encoding/base64"
	"fmt"

	dcfg "summitto.com/txsigner/config/daemon"
)

func GenerateKeyPair() (publicKeyBase64, privateKeyBase64 string, publicKey ed25519.PublicKey, privateKey ed25519.PrivateKey, err error) {
	publicKey, privateKey, err = ed25519.GenerateKey(nil)

	if err != nil {
		fmt.Println("Error:", err)
		return "", "", make([]byte, 0), make([]byte, 0), err
	}

	publicKeyBase64 = base64.StdEncoding.EncodeToString(
		[]byte(publicKey),
	)

	privateKeyBase64 = base64.StdEncoding.EncodeToString(
		[]byte(privateKey),
	)

	return
}

func CalculatePublicKey(privateKeyBase64 string) (publicKeyBase64 string) {
	privateKey, _ := base64.StdEncoding.DecodeString(privateKeyBase64)
	private := ed25519.NewKeyFromSeed(privateKey[:32])
	pub := []byte(private.Public().(ed25519.PublicKey))
	publicKeyBase64 = base64.StdEncoding.EncodeToString(pub)
	return
}

func GetPrivateKey(retrievePrivateKey func() string) (privateKeyBase64 string, privateKey ed25519.PrivateKey, err error) {
	envPrivateKey := retrievePrivateKey()
	var decodedPrivateKey []byte
	decodedPrivateKey, err = base64.StdEncoding.DecodeString(envPrivateKey)
	if err != nil {
		fmt.Println("Error decoding the private key", err)
		return "", nil, err
	}
	return envPrivateKey, ed25519.PrivateKey(decodedPrivateKey), err
}

func GetDaemonPrivateKey() (privateKeyBase64 string, privateKey ed25519.PrivateKey, err error) {
	privateKeyBase64, privateKey, err = GetPrivateKey(func() string {
		return dcfg.GetDaemonConfig().PrivateKey
	})
	return
}

func GetDaemonPublicKey() (publicKeyBase64 string, publicKey ed25519.PublicKey, err error) {
	privateKeyBase64, _, err := GetDaemonPrivateKey()
	publicKeyBase64 = CalculatePublicKey(privateKeyBase64)

	publicKey, err = base64.StdEncoding.DecodeString(publicKeyBase64)
	return
}
