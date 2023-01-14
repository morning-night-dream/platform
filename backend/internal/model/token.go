package model

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

type uidCtxKey struct{}

const (
	idTokenExpire      = time.Minute * 15
	refreshTokenExpire = time.Hour * 24
	tokenSecret        = "secret" // 要環境変数
)

var ErrInvalid = errors.New("invalid")

type IDToken struct {
	jwt.Token
}

func (i IDToken) String() string {
	s, err := i.SignedString([]byte(tokenSecret))
	if err != nil {
		return ""
	}

	return s
}

func (i IDToken) GetUID() string {
	return ""
}

type RefreshToken struct {
	jwt.Token
}

func (r RefreshToken) String() string {
	s, err := r.SignedString([]byte(tokenSecret))
	if err != nil {
		return ""
	}

	return s
}

func CreateTokenFrom(tokenStr string) (jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Printf("unexpected signing method: %v", token.Header["alg"])

			return nil, ErrInvalid
		}

		return []byte(tokenSecret), nil
	})
	if err != nil {
		log.Print(err.Error())

		return jwt.Token{}, ErrInvalid
	}

	return *token, nil
}

type Token struct {
	IDToken      IDToken
	RefreshToken RefreshToken
}

// @reference https://medium.com/monstar-lab-bangladesh-engineering/jwt-auth-in-go-part-2-refresh-tokens-d334777ca8a0
func GenerateToken(userID string) (Token, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return Token{}, ErrInvalid
	}

	claims["uid"] = userID
	claims["sub"] = userID
	claims["exp"] = time.Now().Add(idTokenExpire).Unix()

	refreshToken := jwt.New(jwt.SigningMethodHS256)

	rtClaims, ok := refreshToken.Claims.(jwt.MapClaims)
	if !ok {
		return Token{}, ErrInvalid
	}

	rtClaims["sub"] = userID
	rtClaims["exp"] = time.Now().Add(refreshTokenExpire).Unix()

	return Token{
		IDToken:      IDToken{*token},
		RefreshToken: RefreshToken{*refreshToken},
	}, nil
}
