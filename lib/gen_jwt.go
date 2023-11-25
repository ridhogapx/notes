package lib

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func GenJWT(body interface{}) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": body,
	})

	stringifiyToken, err := token.SignedString("mySecret")

	if err != nil {
		fmt.Println("There is error in encode token:", err)
		return ""
	}

	return stringifiyToken
}
