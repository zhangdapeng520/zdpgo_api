package main

import (
	"net/http"
	"time"

	"github.com/zhangdapeng520/zdpgo_api/gin"
	"github.com/zhangdapeng520/zdpgo_api/gin/binding"
	"github.com/zhangdapeng520/zdpgo_api/validator"
)

// Booking 图书预订，包含绑定数据和校验数据
type Booking struct {
	// 借书时间
	// bookabledate 借书时间校验
	CheckIn time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`

	// 还书时间
	// gtfield=CheckIn 大于借书时间
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

// 自定义校验器
var bookableDate validator.Func = func(fl validator.FieldLevel) bool {
	// 获取字段的接口并转换为时间类型
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		today := time.Now()

		// 如果是当前时间之前的时间，则校验不通过
		if today.After(date) {
			return false
		}
	}
	return true
}

func getBookable(c *gin.Context) {
	var b Booking
	if err := c.ShouldBindWith(&b, binding.Query); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "预订日期有效！"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

// 设置路由
func setupRouter() *gin.Engine {
	// 创建路由
	route := gin.Default()

	// 绑定校验器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册自定义校验器
		v.RegisterValidation("bookabledate", bookableDate)
	}

	// 路径
	route.GET("/bookable", getBookable)

	// 返回
	return route
}

func main() {
	// 创建路由
	route := setupRouter()

	// 监听
	route.Run(":8085")
}
