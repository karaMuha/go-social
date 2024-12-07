package domain

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var PrivateKey *rsa.PrivateKey

func InitPrivateKey(filename string) error {
	file, err := os.Open(filename)

	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)

	_, err = file.Read(buffer)
	if err != nil {
		return err
	}

	data, _ := pem.Decode(buffer)
	key, err := x509.ParsePKCS8PrivateKey(data.Bytes)

	if err != nil {
		return err
	}

	if key, ok := key.(*rsa.PrivateKey); ok {
		PrivateKey = key
		return nil
	}

	return errors.New("error while reading private key")
}

func GenerateJwt(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"sub": userId,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	return token.SignedString(PrivateKey)
}
