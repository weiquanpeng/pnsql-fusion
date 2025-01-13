package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/sync/singleflight"
)

// JWT 定义、生成密钥
type JWT struct {
	SigningKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte("www.PnSql.com"),
	}
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

type CustomClaims struct {
	Id         string `json:"id"`
	Username   string `json:"username"`
	BufferTime int64
	jwt.StandardClaims
}

func NewCustomClaims(id, username string, bufferTime int64, iat, nat, exp int64) CustomClaims {
	return CustomClaims{
		Id:         id,
		Username:   username,
		BufferTime: bufferTime,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  iat,         // 签发时间
			NotBefore: nat,         //生效时间
			ExpiresAt: exp,         // 过期时间
			Audience:  "GaoGao",    //受众
			Issuer:    "Pwq-Admin", //签发者
		},
	}
}

// 创建一个 JWT 令牌
func (t *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(t.SigningKey)
}

// 解析 JWT
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("token无效")
}

// CreateTokenByOldToken 旧token 换新token 使用归并回源避免并发问题
func (j *JWT) CreateTokenByOldToken(oldToken string, claims CustomClaims) (string, error) {
	group := singleflight.Group{}
	v, err, _ := group.Do("JWT:"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})
	return v.(string), err
}
