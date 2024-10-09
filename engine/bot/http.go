package bot

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 跨域 middleware
func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS,PUT,DELETE,UPDATE")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token,Authorization,Token")
		c.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func runServer() {
	r := gin.New()
	r.Use(cors())
	// 消息回调
	r.POST("/wxbot/callback", func(c *gin.Context) {
		bot.framework.Callback(c, eventBuffer.ProcessEvent)
	})
	r.Run("0.0.0.0:8080")
}
