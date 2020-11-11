package transactions

import (
	"crypto/ed25519"
	"encoding/base64"
	"fmt"
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

func GetPrivateKey(privateKeyString string) (privateKeyBase64 string, privateKey ed25519.PrivateKey, err error) {
	var decodedPrivateKey []byte
	decodedPrivateKey, err = base64.StdEncoding.DecodeString(privateKeyString)
	if err != nil {
		fmt.Println("Error decoding the private key", err)
		return "", nil, err
	}
	return privateKeyString, ed25519.PrivateKey(decodedPrivateKey), err
}

func GetDaemonPublicKey(privateKeyBase64 string) (publicKeyBase64 string, publicKey ed25519.PublicKey, err error) {
	publicKeyBase64 = CalculatePublicKey(privateKeyBase64)

	publicKey, err = base64.StdEncoding.DecodeString(publicKeyBase64)
	return
}
