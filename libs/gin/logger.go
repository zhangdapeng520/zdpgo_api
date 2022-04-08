// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gin

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/mattn/go-isatty"
)

type consoleColorModeValue int

const (
	autoColor consoleColorModeValue = iota
	disableColor
	forceColor
)

const (
	green   = "\033[97;42m"
	white   = "\033[90;47m"
	yellow  = "\033[90;43m"
	red     = "\033[97;41m"
	blue    = "\033[97;44m"
	magenta = "\033[97;45m"
	cyan    = "\033[97;46m"
	reset   = "\033[0m"
)

var consoleColorMode = autoColor // 控制台日志

// LoggerConfig 日志中间件的配置
type LoggerConfig struct {
	Formatter      LogFormatter // 日志格式化，可选的，默认是gin.defaultLogFormatter
	Output         io.Writer    // 日志输出流，可选的，默认是gin.DefaultWriter
	SkipPaths      []string     // 不记录日志的路径，可选参数
	IsDetailLogger bool         // 是否使用详细日志
}

// LogFormatter 日志格式化的函数签名
type LogFormatter func(params LogFormatterParams) string

// LogFormatterParams 当记录时间到来时，任何格式化程序都会被交给这个结构
type LogFormatterParams struct {
	Request *http.Request

	// TimeStamp 显示服务器返回响应后的时间。
	TimeStamp time.Time

	// StatusCode HTTP状态码
	StatusCode int

	// Latency 消耗时间
	Latency time.Duration

	// ClientIP 等于Context's ClientIP方法
	ClientIP string

	// Method HTTP请求方法
	Method string

	// Path 客户端请求路径
	Path string

	// ErrorMessage 错误消息的集合
	ErrorMessage string

	// isTerm 显示gin的输出描述符是否引用终端。
	isTerm bool

	// BodySize 是响应体的大小
	BodySize int

	// Keys 是根据请求的上下文设置的键。
	Keys map[string]interface{}

	// 是否显示请求头
	IsHeader bool
}

// StatusCodeColor ANSI颜色，用于将http状态代码正确记录到终端。
func (p *LogFormatterParams) StatusCodeColor() string {
	code := p.StatusCode
	switch {
	case code >= http.StatusOK && code < http.StatusMultipleChoices:
		return green
	case code >= http.StatusMultipleChoices && code < http.StatusBadRequest:
		return white
	case code >= http.StatusBadRequest && code < http.StatusInternalServerError:
		return yellow
	default:
		return red
	}
}

// MethodColor ANSI颜色，用于将http方法正确记录到终端。
func (p *LogFormatterParams) MethodColor() string {
	method := p.Method
	switch method {
	case http.MethodGet:
		return blue
	case http.MethodPost:
		return cyan
	case http.MethodPut:
		return yellow
	case http.MethodDelete:
		return red
	case http.MethodPatch:
		return green
	case http.MethodHead:
		return magenta
	case http.MethodOptions:
		return white
	default:
		return reset
	}
}

// ResetColor 重置所有转义属性。
func (p *LogFormatterParams) ResetColor() string {
	return reset
}

// IsOutputColor 指示是否可以将颜色输出到日志。
func (p *LogFormatterParams) IsOutputColor() bool {
	return consoleColorMode == forceColor || (consoleColorMode == autoColor && p.isTerm)
}

// defaultLogFormatter 是记录器中间件使用的默认日志格式函数。
var defaultLogFormatter = func(param LogFormatterParams) string {
	var statusColor, methodColor, resetColor string
	if param.IsOutputColor() {
		statusColor = param.StatusCodeColor()
		methodColor = param.MethodColor()
		resetColor = param.ResetColor()
	}

	if param.Latency > time.Minute {
		// Truncate in a golang < 1.8 safe way
		param.Latency = param.Latency - param.Latency%time.Second
	}

	// 默认格式
	return fmt.Sprintf("[ZDPGO_API] %v |%s %3d %s| %13v | %15s |%s %-7s %s %#v\n%s",
		param.TimeStamp.Format("2006-01-02 15:04:05"), // 日期时间
		statusColor, param.StatusCode, resetColor, // 状态码
		param.Latency,                         // 消耗时间
		param.ClientIP,                        // 客户端IP
		methodColor, param.Method, resetColor, // HTTP方法
		param.Path,         // 请求路径
		param.ErrorMessage, // 错误消息
	)
}

// detailLogFormatter 记录详细的日志信息
var detailLogFormatter = func(param LogFormatterParams) string {
	var statusColor, methodColor, resetColor string
	if param.IsOutputColor() {
		statusColor = param.StatusCodeColor()
		methodColor = param.MethodColor()
		resetColor = param.ResetColor()
	}

	if param.Latency > time.Minute {
		// Truncate in a golang < 1.8 safe way
		param.Latency = param.Latency - param.Latency%time.Second
	}

	// 详细日志格式
	jsonHeaderData, _ := json.Marshal(param.Request.Header)
	jsonFormData, _ := json.Marshal(param.Request.PostForm)
	defaultFormat := fmt.Sprintf("[ZDPGO_API] %v |%s %3d %s| %13v | %15s |%s %-7s %s %#v\n%s \n请求头信息：%s \nForm表单信息：%s \nBody主体信息：%s\n",
		param.TimeStamp.Format("2006-01-02 15:04:05"), // 日期时间
		statusColor, param.StatusCode, resetColor, // 状态码
		param.Latency,                         // 消耗时间
		param.ClientIP,                        // 客户端IP
		methodColor, param.Method, resetColor, // HTTP方法
		param.Path,         // 请求路径
		param.ErrorMessage, // 错误消息
		jsonHeaderData,     // 请求头信息
		jsonFormData,       // 表单信息
		param.Request.Body, // 主体信息
	)

	return defaultFormat
}

// DisableConsoleColor disables color output in the console.
func DisableConsoleColor() {
	consoleColorMode = disableColor
}

// ForceConsoleColor force color output in the console.
func ForceConsoleColor() {
	consoleColorMode = forceColor
}

// ErrorLogger returns a handlerfunc for any error type.
func ErrorLogger() HandlerFunc {
	return ErrorLoggerT(ErrorTypeAny)
}

// ErrorLoggerT returns a handlerfunc for a given error type.
func ErrorLoggerT(typ ErrorType) HandlerFunc {
	return func(c *Context) {
		c.Next()
		errors := c.Errors.ByType(typ)
		if len(errors) > 0 {
			c.JSON(-1, errors)
		}
	}
}

// Logger instances a Logger middleware that will write the logs to gin.DefaultWriter.
// By default gin.DefaultWriter = os.Stdout.
func Logger() HandlerFunc {
	return LoggerWithConfig(LoggerConfig{})
}

// LoggerWithFormatter instance a Logger middleware with the specified log format function.
func LoggerWithFormatter(f LogFormatter) HandlerFunc {
	return LoggerWithConfig(LoggerConfig{
		Formatter: f,
	})
}

// LoggerWithWriter instance a Logger middleware with the specified writer buffer.
// Example: os.Stdout, a file opened in write mode, a socket...
func LoggerWithWriter(out io.Writer, notlogged ...string) HandlerFunc {
	return LoggerWithConfig(LoggerConfig{
		Output:    out,
		SkipPaths: notlogged,
	})
}

// LoggerWithConfig 根据配置返回一个logger日志中间件
func LoggerWithConfig(conf LoggerConfig) HandlerFunc {
	// 格式化
	formatter := conf.Formatter
	if formatter == nil {
		if conf.IsDetailLogger {
			formatter = detailLogFormatter // 详细日志
		} else {
			formatter = defaultLogFormatter // 默认日志
		}
	}

	// 输出
	out := conf.Output
	if out == nil {
		out = DefaultWriter
	}

	// 不记录
	notlogged := conf.SkipPaths

	isTerm := true

	if w, ok := out.(*os.File); !ok || os.Getenv("TERM") == "dumb" ||
		(!isatty.IsTerminal(w.Fd()) && !isatty.IsCygwinTerminal(w.Fd())) {
		isTerm = false
	}

	var skip map[string]struct{}

	if length := len(notlogged); length > 0 {
		skip = make(map[string]struct{}, length)

		for _, path := range notlogged {
			skip[path] = struct{}{}
		}
	}

	return func(c *Context) {
		// 开始时间
		start := time.Now()
		path := c.Request.URL.Path    // 路径
		raw := c.Request.URL.RawQuery // 查询参数

		// 通过请求
		c.Next()

		// 只记录不跳过的路径
		if _, ok := skip[path]; !ok {
			param := LogFormatterParams{
				Request: c.Request,
				isTerm:  isTerm,
				Keys:    c.Keys,
			}

			// 停止计时
			param.TimeStamp = time.Now()
			param.Latency = param.TimeStamp.Sub(start)
			param.ClientIP = c.ClientIP()
			param.Method = c.Request.Method
			param.StatusCode = c.Writer.Status()
			param.ErrorMessage = c.Errors.ByType(ErrorTypePrivate).String()
			param.BodySize = c.Writer.Size()

			if raw != "" {
				path = path + "?" + raw
			}

			param.Path = path

			// 输出日志
			fmt.Fprint(out, formatter(param))
		}
	}
}
