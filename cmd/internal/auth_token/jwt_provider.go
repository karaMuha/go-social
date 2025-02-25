package authtoken

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
	"path/filepath"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtProvider struct {
	privateKey *rsa.PrivateKey
}

var _ ITokenProvider = (*JwtProvider)(nil)

func NewTokenProvider(pathToPrivateKey string) JwtProvider {
	return JwtProvider{
		privateKey: initPrivateKey(pathToPrivateKey),
	}
}

func initPrivateKey(filename string) *rsa.PrivateKey {
	file, err := os.Open(filepath.Clean(filename))

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

func (p JwtProvider) GenerateToken(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"sub": userId,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	return token.SignedString(p.privateKey)
}

func (p JwtProvider) VerifyToken(token string) (any, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodRSA)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return p.privateKey.Public(), nil
	})

	if err != nil {
		return nil, err
	}

	if !parsedToken.Valid {
		return nil, errors.New("invalid token")
	}

	return parsedToken, nil
}

func (p JwtProvider) GetUserIDFromToken(parsedToken any) (string, error) {
	jwtToken, ok := parsedToken.(*jwt.Token)
	if !ok {
		return "", errors.New("could not convert token to jwt")
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)

	if !ok {
		return "", errors.New("could not convert jwt claims")
	}

	userID, ok := claims["sub"].(string)

	if !ok {
		return "", errors.New("could not convert user id from jwt claims to string")
	}

	return userID, nil
}
