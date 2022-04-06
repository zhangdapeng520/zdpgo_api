package main

import (
	"net/http"
	"time"

	zdp_validator "zdpgo_gin/validator"

	"github.com/zhangdapeng520/zdpgo_gin/libs/gin"
	"github.com/zhangdapeng520/zdpgo_gin/libs/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Booking contains binded and validated data.
type Booking struct {
	Mobile string    `form:"mobile" binding:"required,mobile"`
	Date   time.Time `form:"date" binding:"required" time_format:"2006-01-02"`
}

func main() {
	route := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("mobile", zdp_validator.ValidateMobile)
	}

	// http://127.0.0.1:8080/bookable?mobile=18811118888&date=2022-01-01
	route.GET("/bookable", getBookable)
	route.Run(":8080")
}

func getBookable(c *gin.Context) {
	var b Booking
	if err := c.ShouldBindWith(&b, binding.Query); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
