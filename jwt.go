package zdpgo_gin

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Claims jwt校验对象
type Claims struct {
	ID       uint32                 `json:"id"`       // id
	Username string                 `json:"username"` // 用户名
	Email    string                 `json:"email"`    // 邮箱
	Mobile   string                 `json:"mobile"`   // 手机号
	Role     uint                   `json:"role"`     // 角色
	Message  string                 `json:"message"`  // 要传递的消息
	Data     map[string]interface{} `json:"data"`     // 要传递的数据
	jwt.StandardClaims
}

var (
	TokenExpired     = errors.New("Token已过期")
	TokenNotValidYet = errors.New("Token未激活")
	TokenMalformed   = errors.New("Token格式错误")
	TokenInvalid     = errors.New("Token无效")
)

// CreateToken 创建一个token
// @param claims 要创建token的数据
// @return 成功返回token,失败返回错误信息
func (g *Gin) CreateToken(claims Claims) (string, error) {
	// 过期时间默认3小时
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(g.config.Jwt.JwtExpired)).Unix()

	// 创建token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 返回token
	g.log.Info("使用的key是什么", "key", g.config.Jwt.JwtKey)
	return token.SignedString([]byte(g.config.Jwt.JwtKey))
}

// ParseToken 解析 token
// @param tokenString token字符串
// @return 成功返回解析结果,失败返回错误信息
func (g *Gin) ParseToken(tokenString string) (*Claims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(g.config.Jwt.JwtKey), nil
	})
	g.log.Info("解析token", "token", token, "error", err)

	// 解析失败
	if err != nil {
		g.log.Error("解析失败", "error", err.Error())
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}

	// 解析成功
	if token != nil {
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			g.log.Info("解析token成功", "claims", claims)
			return claims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}

}

// RefreshToken 更新token
// @param tokenString 之前的token
// @return 成功返回新的token,失败返回错误信息
func (g *Gin) RefreshToken(tokenString string) (string, error) {
	// 当前时间
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}

	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return g.config.Jwt.JwtKey, nil
	})
	if err != nil {
		return "", err
	}

	// 刷新token
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(3 * time.Hour).Unix()
		return g.CreateToken(*claims)
	}
	return "", TokenInvalid
}
