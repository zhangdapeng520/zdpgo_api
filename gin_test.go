package zdpgo_gin

import (
	"fmt"
	"testing"
)

// 准备gin对象
func prepareGin() *Gin {
	g := New(GinConfig{
		Debug: true,
	})
	return g
}

// 测试新建gin对象
func TestGin_New(t *testing.T) {
	g := prepareGin()
	fmt.Println(g)
}
