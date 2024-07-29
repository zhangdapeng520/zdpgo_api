package middleware

import (
	"bytes"
	"context"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"
)

var (
	// LogEntryCtxKey 是用于存储请求日志条目的 context.Context 密钥。
	LogEntryCtxKey = &contextKey{"LogEntry"}

	// DefaultLogger 日志中间件处理程序会调用来记录每个请求。
	// DefaultLogger 是由 Logger 中间件处理程序调用的，用于记录每个请求。
	DefaultLogger func(next http.Handler) http.Handler
)

func init() {
	color := true
	if runtime.GOOS == "windows" {
		color = false
	}
	DefaultLogger = RequestLogger(&DefaultLogFormatter{Logger: log.New(os.Stdout, "", log.LstdFlags), NoColor: !color})
}

// Logger 是一个中间件，它记录每次请求的开始和结束，以及一些有用的数据，包括请求内容、响应状态和返回时间。
// 当标准输出为 TTY 时，日志程序将以彩色打印，否则将以黑白打印。如果提供了请求 ID，日志程序将打印该 ID。
//
// 此外，还可以访问 https://github.com/goware/httplog，了解更深入的 http 日志记录器，并支持结构化日志记录。
//
// 重要注意事项：记录仪应在任何可能更改响应的其他中间件（如 middleware.Recoverer）之前运行。示例
//
//	r := api.NewRouter()
//	r.Use(middleware.Logger)        // <--<< Logger should come before Recoverer
//	r.Use(middleware.Recoverer)
//	r.Get("/", handler)
func Logger(next http.Handler) http.Handler {
	return DefaultLogger(next)
}

// RequestLogger 返回使用自定义 LogFormatter 的日志处理程序。
func RequestLogger(f LogFormatter) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			entry := f.NewLogEntry(r)
			ww := NewWrapResponseWriter(w, r.ProtoMajor)

			t1 := time.Now()
			defer func() {
				entry.Write(ww.Status(), ww.BytesWritten(), ww.Header(), time.Since(t1), nil)
			}()

			next.ServeHTTP(ww, WithLogEntry(r, entry))
		}
		return http.HandlerFunc(fn)
	}
}

// LogFormatter initiates the beginning of a new LogEntry per request.
// See DefaultLogFormatter for an example implementation.
type LogFormatter interface {
	NewLogEntry(r *http.Request) LogEntry
}

// LogEntry records the final log when a request completes.
// See defaultLogEntry for an example implementation.
type LogEntry interface {
	Write(status, bytes int, header http.Header, elapsed time.Duration, extra interface{})
	Panic(v interface{}, stack []byte)
}

// GetLogEntry returns the in-context LogEntry for a request.
func GetLogEntry(r *http.Request) LogEntry {
	entry, _ := r.Context().Value(LogEntryCtxKey).(LogEntry)
	return entry
}

// WithLogEntry sets the in-context LogEntry for a request.
func WithLogEntry(r *http.Request, entry LogEntry) *http.Request {
	r = r.WithContext(context.WithValue(r.Context(), LogEntryCtxKey, entry))
	return r
}

// LoggerInterface accepts printing to stdlib logger or compatible logger.
type LoggerInterface interface {
	Print(v ...interface{})
}

// DefaultLogFormatter is a simple logger that implements a LogFormatter.
type DefaultLogFormatter struct {
	Logger  LoggerInterface
	NoColor bool
}

// NewLogEntry creates a new LogEntry for the request.
func (l *DefaultLogFormatter) NewLogEntry(r *http.Request) LogEntry {
	useColor := !l.NoColor
	entry := &defaultLogEntry{
		DefaultLogFormatter: l,
		request:             r,
		buf:                 &bytes.Buffer{},
		useColor:            useColor,
	}

	reqID := GetReqID(r.Context())
	if reqID != "" {
		cW(entry.buf, useColor, nYellow, "[%s] ", reqID)
	}
	cW(entry.buf, useColor, nCyan, "\"")
	cW(entry.buf, useColor, bMagenta, "%s ", r.Method)

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	cW(entry.buf, useColor, nCyan, "%s://%s%s %s\" ", scheme, r.Host, r.RequestURI, r.Proto)

	entry.buf.WriteString("from ")
	entry.buf.WriteString(r.RemoteAddr)
	entry.buf.WriteString(" - ")

	return entry
}

type defaultLogEntry struct {
	*DefaultLogFormatter
	request  *http.Request
	buf      *bytes.Buffer
	useColor bool
}

func (l *defaultLogEntry) Write(status, bytes int, header http.Header, elapsed time.Duration, extra interface{}) {
	switch {
	case status < 200:
		cW(l.buf, l.useColor, bBlue, "%03d", status)
	case status < 300:
		cW(l.buf, l.useColor, bGreen, "%03d", status)
	case status < 400:
		cW(l.buf, l.useColor, bCyan, "%03d", status)
	case status < 500:
		cW(l.buf, l.useColor, bYellow, "%03d", status)
	default:
		cW(l.buf, l.useColor, bRed, "%03d", status)
	}

	cW(l.buf, l.useColor, bBlue, " %dB", bytes)

	l.buf.WriteString(" in ")
	if elapsed < 500*time.Millisecond {
		cW(l.buf, l.useColor, nGreen, "%s", elapsed)
	} else if elapsed < 5*time.Second {
		cW(l.buf, l.useColor, nYellow, "%s", elapsed)
	} else {
		cW(l.buf, l.useColor, nRed, "%s", elapsed)
	}

	l.Logger.Print(l.buf.String())
}

func (l *defaultLogEntry) Panic(v interface{}, stack []byte) {
	PrintPrettyStack(v)
}
