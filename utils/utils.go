package utils

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Real-Musafir/bookshop/config"
	"github.com/golang-jwt/jwt"
)



func MapToStruct(data map[string]interface{}, result interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonData, result)
	if err != nil {
		return err
	}

	return nil
}

func CreateToken(tokenPayload map[string]any) (string, error) {
	secretKey := []byte(config.GetEnvProperty("secret_key"))

	claims := jwt.MapClaims{}
	claims["user_id"] = tokenPayload["user_id"]
	claims["user_name"] = tokenPayload["user_name"]
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return signedToken, nil

}

func VerifyToken(tokenStirng string) (*jwt.MapClaims, error) {
	secretKey := []byte(config.GetEnvProperty("secret_key"))
	token, err := jwt.Parse(tokenStirng, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}

		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, _ := token.Claims.(jwt.MapClaims)

	return &claims, nil
}