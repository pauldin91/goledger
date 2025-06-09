package utils

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"log"
	"math/big"
)

type KeyPair struct {
	keyPair *ecdsa.PrivateKey
	private string
	public  string
}

func NewKeyPair() KeyPair {
	var newKeyPair = KeyPair{}
	pair := genKeyPair(&newKeyPair.private, &newKeyPair.public)
	newKeyPair.keyPair = pair
	return newKeyPair
}

func (pair *KeyPair) GetPublicKey() string {
	return pair.public
}

func (pair *KeyPair) Sign(hashedData string) string {

	res, err := pair.keyPair.Sign(rand.Reader, []byte(hashedData), crypto.SHA256)
	if err != nil {
		log.Fatal("could sign data")
	}
	return string(res)
}

func VerifySignature(publicKey string, hashed []byte, sig []byte) bool {
	keyBytes, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return false
	}

	if len(keyBytes) != 64 {
		return false
	}

	x := new(big.Int).SetBytes(keyBytes[:32])
	y := new(big.Int).SetBytes(keyBytes[32:])

	pubKey := &ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     x,
		Y:     y,
	}

	valid := ecdsa.VerifyASN1(pubKey, hashed, sig)
	return valid
}

func genKeyPair(priv, pub *string) *ecdsa.PrivateKey {
	privKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	privKeyBytes := privKey.D.Bytes()
	*priv = base64.StdEncoding.EncodeToString(privKeyBytes)

	pubKeyBytes := append(privKey.PublicKey.X.Bytes(), privKey.PublicKey.Y.Bytes()...)
	*pub = base64.StdEncoding.EncodeToString(pubKeyBytes)

	return privKey
}

func Hash(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))

}
