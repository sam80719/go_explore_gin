package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		host := c.Request.Host
		url := c.Request.URL
		t := time.Now()
		formatted := fmt.Sprintf(fmt.Sprintf("%d/%02d/%02d %02d:%02d:%02d",
			t.Year(), t.Month(), t.Day(),
			t.Hour(), t.Minute(), t.Second()))
		fmt.Printf("%s: host=%s, url=%s \n", formatted, host, url)
		c.Next()
	}
}
