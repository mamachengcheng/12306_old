package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mamachengcheng/12306/app/resource"
	"github.com/mamachengcheng/12306/app/utils"
	"gopkg.in/ini.v1"
	"log"
	"time"
)

var jwtSecret = []byte(getSignalKey())

func getSignalKey() string {
	cfg, _ := ini.Load(resource.ConfFilePath)
	return cfg.Section("server").Key("sign_key").String()
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(username string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(12 * time.Hour)

	claims := Claims{
		username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "12306",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

func JWTMiddleware() gin.HandlerFunc {
	response := utils.Response{
		Code: 401,
		Data: make(map[string]interface{}),
		Msg:  "未登录",
	}

	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		claims, err := ParseToken(token)
		if err == nil {
			c.Set("claims", claims)
			c.Next()
		} else {
			jwtAbort(response, c)
		}
	}
}

func jwtAbort(response utils.Response, c *gin.Context) {
	utils.StatusOKResponse(response, c)
	c.Abort()
}
