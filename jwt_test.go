package zdpgo_gin

import (
	"fmt"
	"testing"
)

func TestGin_CreateToken(t *testing.T) {
	g := prepareGin()

	// 创建token
	c := Claims{
		ID:       1,
		Username: "zhangdapeng",
		Role:     1,
	}
	token, err := g.CreateToken(c)
	fmt.Println(token, err)

	// 解析token
	claims, err := g.ParseToken(token)
	fmt.Println(claims, err)
}
