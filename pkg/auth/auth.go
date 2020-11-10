package auth

import (
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

var JwtAuth = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//TODO: remove api/game/get and properly pass the authorization header through the iframe in the frontend
		notAuth := []string{"/API/auth", "/static/js/"}
		requestPath := r.URL.Path

		for _, v := range notAuth {
			if v == requestPath || strings.HasPrefix(requestPath, v) {
				fmt.Println(requestPath)
				next.ServeHTTP(w, r)
				return
			}
		}

		_, err := ValidateToken(r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)

	})
}

func ExtractToken(r *http.Request) (string, error) {
	tokenHeader := r.Header.Get("Authorization")

	if tokenHeader == "" {
		fmt.Println(tokenHeader)
		return "", ErrTokenMissing
	}

	splitted := strings.Split(tokenHeader, " ")
	if len(splitted) != 2 {
		return "", ErrTokenMalformed
	}

	tokenPart := splitted[1]
	return tokenPart, nil
}

func ExtractTokenInfo(strToken string) (*Token, error) {
	tk := &Token{}
	token, err := jwt.ParseWithClaims(strToken, tk, func(token *jwt.Token) (interface{}, error) {
		return []byte("tokenPassword"), nil
	})

	if err != nil {
		return tk, ErrTokenMalformed
	}

	if !token.Valid {
		return tk, ErrTokenInvalid
	}

	return tk, nil
}

func ValidateToken(r *http.Request) (*Token, error) {
	strToken, _err := ExtractToken(r)

	if _err != nil {
		return nil, _err
	}

	token, err := ExtractTokenInfo(strToken)

	return token, err

}
