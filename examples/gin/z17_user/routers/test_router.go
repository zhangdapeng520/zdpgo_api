package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"zdpgo_gin/examples/z17_user/globals"
)

func testLog(context *gin.Context) {
	//Info级别的日志
	globals.Logrus.WithFields(logrus.Fields{
		"name": "hanyun",
	}).Info("记录一下Info日志", "Info")

	//Error级别的日志
	globals.Logrus.WithFields(logrus.Fields{
		"name": "hanyun",
	}).Error("记录一下Error日志", "Error")

	//Warn级别的日志
	globals.Logrus.WithFields(logrus.Fields{
		"name": "hanyun",
	}).Warn("记录一下Warn日志", "Warn")

	//Debug级别的日志
	globals.Logrus.WithFields(logrus.Fields{
		"name": "hanyun",
	}).Debug("记录一下Debug日志", "Debug")
}
