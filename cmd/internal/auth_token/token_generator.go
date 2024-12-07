package authtoken

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenGenerator struct {
	privateKey *rsa.PrivateKey
}

func NewTokenGenerator(pathToPrivateKey string) TokenGenerator {
	return TokenGenerator{
		privateKey: initPrivateKey(pathToPrivateKey),
	}
}

func initPrivateKey(filename string) *rsa.PrivateKey {
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)

	_, err = file.Read(buffer)
	if err != nil {
		panic(err)
	}

	data, _ := pem.Decode(buffer)
	key, err := x509.ParsePKCS8PrivateKey(data.Bytes)

	if err != nil {
		panic(err)
	}

	if key, ok := key.(*rsa.PrivateKey); ok {
		return key

	}

	panic("error while reading private key, cannot convert private key")
}

func (t TokenGenerator) GenerateToken(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"sub": userId,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	return token.SignedString(t.privateKey)
}

func (t TokenGenerator) VerifyToken(jwtToken string) (*jwt.Token, error) {
	parsedToken, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodRSA)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return t.privateKey.Public(), nil
	})

	if err != nil {
		return nil, err
	}

	if !parsedToken.Valid {
		return nil, errors.New("invalid token")
	}

	return parsedToken, nil
}
