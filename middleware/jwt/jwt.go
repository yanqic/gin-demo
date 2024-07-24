package jwt

import (
	"errors"
	"fmt"
	"gin-demo/config"
	"gin-demo/pkg/util"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data interface{}

		code := http.StatusOK
		errMsg := ""
		Authorization := c.GetHeader("Authorization")
		token := strings.Split(Authorization, " ")

		if Authorization == "" {
			code = http.StatusBadRequest
		} else {
			claims, err := j.ParseToken(token[1])
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = http.StatusUnauthorized
					errMsg = "授权已过期"
				default:
					code = http.StatusInternalServerError
					errMsg = err.Error()
				}
			}
			c.Set("claims", claims)
		}

		if code != http.StatusOK {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  errMsg,
				"data": data,
			})

			c.Abort()
			return
		}
		c.Next()
	}
}

// 一些常量
var (
	SignKey = config.JwtConfig.Secret
	Timeout = config.JwtConfig.Timeout
	j       = NewJWT()
)

// JWT 签名结构
type JWT struct {
	SigningKey []byte
}

// 新建一个jwt实例
func NewJWT() *JWT {
	return &JWT{
		[]byte(SignKey),
	}
}

type Claims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func (j *JWT) GenerateToken(id int, username, password string, nowTime time.Time) (string, error) {
	expireTime := nowTime.Add(time.Duration(Timeout))
	fmt.Print(id)
	claims := Claims{
		id,
		username,
		util.EncodeMD5(password),
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "yanqi",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(j.SigningKey)
	return token, err
}

func (j *JWT) ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	if token != nil {
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			return claims, nil
		}
	}

	return nil, err
}

// 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return j.GenerateToken(claims.Id, claims.Username, claims.Password, time.Now())
	}
	return "", errors.New("refresh token faild")
}
