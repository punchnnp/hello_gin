package middleware

import (
	"fmt"
	"gin/handler"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthoruzationMiddleware(c *gin.Context) {
	s := c.Request.Header.Get("Authorization")

	fmt.Println(s)

	ss := strings.Split(s, " ")
	switch a := ss[0]; a {
	case "Bearer":
		b := ss[1]
		err := handler.ValidateToken(b)
		if err != nil {
			fmt.Println(err)
			c.String(http.StatusUnauthorized, err.Error())
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	default:
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}
