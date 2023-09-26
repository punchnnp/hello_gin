package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type UserClaim struct {
	jwt.RegisteredClaims
	Name      string
	ExpiresAt int
}

func LoginHandler(c *gin.Context) {
	signature := []byte("flowers")
	id := c.Param("name")
	// expire should be time format unix for any timezone
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim{
		Name: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(5 * time.Minute)},
		},
	})

	ss, err := token.SignedString(signature)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"token": ss,
	})
}

func ValidateToken(token string) error {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte("flowers"), nil
	})
	if err != nil {
		fmt.Println(err)
		return err
	}

	// check if token valid
	// _, ok := newToken.Claims.(jwt.Claims)
	// if !ok || !newToken.Valid {

	// 	return errors.New("token invalid")
	// }

	// check if expire or not by numeric time
	// time, err := newToken.Claims.GetExpirationTime()
	// fmt.Println(time)

	return nil
}
