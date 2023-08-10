package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"ruoyi-go/config"
	"strings"
	"time"
)

func CreateToken(UserName string, UserId int, DeptId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_name": UserName, // 为了添加createBy 字段
		"user_id":   UserId,
		"dept_id":   DeptId,
		"exp":       time.Now().Unix() + int64(config.Jwt.JwtTtl),
		"iss":       "Ruoyi-Go",
	})

	mySigningKey := []byte(config.Jwt.Secret)

	return token.SignedString(mySigningKey)
}

func VerifyToken(tokenStr string) (*jwt.Token, error) {
	mySigningKey := []byte(config.Jwt.Secret)
	tokenStr = strings.ReplaceAll(tokenStr, config.HeaderSignTokenStr, "")
	return jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
}

func JWTAuthMiddleware() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// 根据实际情况取TOKEN, 这里从request header取
		header := ctx.Request.Header
		tokenStr := header.Get(config.HeaderSignToken)
		if len(tokenStr) < 1 {
			ctx.JSON(http.StatusOK, gin.H{
				"msg":  "参数错误",
				"code": http.StatusInternalServerError,
			})
			ctx.Abort()
			return
		}

		token, err := VerifyToken(tokenStr)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"msg":  "认证失败",
				"code": http.StatusUnauthorized,
			})
			ctx.Abort()
			return
		}
		userId := token.Claims.(jwt.MapClaims)["user_id"]
		deptId := token.Claims.(jwt.MapClaims)["dept_id"]
		userName := token.Claims.(jwt.MapClaims)["user_name"]
		// 此处已经通过了, 可以把Claims中的有效信息拿出来放入上下文使用
		ctx.Set("userId", userId)
		ctx.Set("deptId", deptId)
		ctx.Set("userName", userName)
		ctx.Next()
	}
}
