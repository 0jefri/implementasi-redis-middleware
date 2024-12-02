package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/project-sistem-voucher/config"
	"github.com/project-sistem-voucher/helper"
)

type Middleware struct {
	Cacher config.Cacher
}

func NewMiddleware(cacher config.Cacher) Middleware {
	return Middleware{
		Cacher: cacher,
	}
}

func (m *Middleware) Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		id := c.GetHeader("ID-KEY")
		val, err := m.Cacher.Get(id)
		if err != nil {
			helper.BadResponse(c, err.Error(), http.StatusInternalServerError)
			c.Abort()
			return
		}

		if val == "" || val != token {
			helper.BadResponse(c, "Unauthorized", http.StatusUnauthorized)
			c.Abort()
			return
		}

		// before request
		c.Next()

	}
}
