package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sakura/record-api/api"
	"github.com/sakura/record-api/utils"
)

// JWTAuth JWT认证中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从 HTTP 头中获取 token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, api.NewResponse(http.StatusUnauthorized, "未提供认证令牌", nil))
			c.Abort()
			return
		}

		// 检查 Authorization 格式
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, api.NewResponse(http.StatusUnauthorized, "认证令牌格式错误", nil))
			c.Abort()
			return
		}

		// 解析 token
		token := parts[1]
		claims, err := utils.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, api.NewResponse(http.StatusUnauthorized, "无效的认证令牌: "+err.Error(), nil))
			c.Abort()
			return
		}

		// 将用户信息保存到上下文
		c.Set("userId", claims.UserID)
		c.Set("userClaims", claims)
		c.Next()
	}
}

// GetUserIDFromContext 从上下文中获取用户ID
func GetUserIDFromContext(c *gin.Context) (uint, bool) {
	userID, exists := c.Get("userId")
	if !exists {
		return 0, false
	}
	return userID.(uint), true
}

// GetUserClaimsFromContext 从上下文中获取用户声明
func GetUserClaimsFromContext(c *gin.Context) (*utils.JWTClaims, bool) {
	claims, exists := c.Get("userClaims")
	if !exists {
		return nil, false
	}
	return claims.(*utils.JWTClaims), true
}
