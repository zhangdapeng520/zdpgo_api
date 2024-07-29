package cors

// cors 软件包是 net/http 处理程序，用于处理 http://www.w3.org/TR/cors/ 所定义的 CORS 相关请求。
//
// 您可以通过向 cors.New 传递一个选项结构来配置它：
//
//     c := cors.New(cors.Options{
//         AllowedOrigins: []string{"foo.com"},
//         AllowedMethods: []string{"GET", "POST", "DELETE"},
//         AllowCredentials: true,
//     })
//
// 然后在链中插入处理程序：
//
//     handler = c.Handler(handler)
//
// 更多选项请参见选项文档。
//
// 由此产生的处理程序是一个标准的 net/http 处理程序。

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// Options 是一个配置容器，用于设置 CORS 中间件。
type Options struct {
	// AllowedOrigins 是跨域请求可执行的来源列表。
	// 如果列表中存在特殊的 "*"值，则允许所有来源。
	// 起源可以包含通配符 (*) 以替换 0 个或多个字符。
	// 例如：http://*.domain.com）。使用通配符会对性能造成一定影响。
	// 每个原点只能使用一个通配符。
	// 默认值为["*"]
	AllowedOrigins []string

	// AllowOriginFunc 是一个用于验证来源的自定义函数。
	// 它以起源为参数，如果允许则返回 true，否则返回 false。
	// 如果设置了该选项，AllowedOrigins 的内容将被忽略。
	AllowOriginFunc func(r *http.Request, origin string) bool

	// AllowedMethods 是允许客户端在跨域请求中使用的方法列表。
	// 默认值为简单方法（HEAD、GET 和 POST）。
	AllowedMethods []string

	// AllowedHeaders 是允许客户端在跨域请求中使用的非简单标头列表。
	// 如果列表中存在特殊的 "*"值，则允许使用所有头信息。
	// 默认值为[]，但 "Origin "总是附加到列表中。
	AllowedHeaders []string

	// ExposedHeaders 表示哪些头可以安全地暴露给 CORS API 规范的 API
	ExposedHeaders []string

	// AllowCredentials 表示请求是否可以包含用户凭证，如 cookie、HTTP 身份验证或客户端 SSL 证书。
	AllowCredentials bool

	// MaxAge 表示预检请求的结果可以缓存多长时间（秒）。
	MaxAge int

	// OptionsPassthrough 指示 prelight 让其他潜在的下一个处理程序处理 OPTIONS 方法。
	// 如果您的应用程序处理 OPTIONS，请将其打开。
	OptionsPassthrough bool

	// Debug 标记增加额外输出，以调试服务器端 CORS 问题
	Debug bool
}

// Logger 日志通用接口
type Logger interface {
	Printf(string, ...interface{})
}

// Cors http 处理程序
type Cors struct {
	Log Logger

	// Normalized list of plain allowed origins
	allowedOrigins []string

	// List of allowed origins containing wildcards
	allowedWOrigins []wildcard

	// Optional origin validator function
	allowOriginFunc func(r *http.Request, origin string) bool

	// Normalized list of allowed headers
	allowedHeaders []string

	// Normalized list of allowed methods
	allowedMethods []string

	// Normalized list of exposed headers
	exposedHeaders []string
	maxAge         int

	// Set to true when allowed origins contains a "*"
	allowedOriginsAll bool

	// Set to true when allowed headers contains a "*"
	allowedHeadersAll bool

	allowCredentials  bool
	optionPassthrough bool
}

// New creates a new Cors handler with the provided options.
func New(options Options) *Cors {
	c := &Cors{
		exposedHeaders:    convert(options.ExposedHeaders, http.CanonicalHeaderKey),
		allowOriginFunc:   options.AllowOriginFunc,
		allowCredentials:  options.AllowCredentials,
		maxAge:            options.MaxAge,
		optionPassthrough: options.OptionsPassthrough,
	}
	if options.Debug && c.Log == nil {
		c.Log = log.New(os.Stdout, "[cors] ", log.LstdFlags)
	}

	// Normalize options
	// Note: for origins and methods matching, the spec requires a case-sensitive matching.
	// As it may error prone, we chose to ignore the spec here.

	// Allowed Origins
	if len(options.AllowedOrigins) == 0 {
		if options.AllowOriginFunc == nil {
			// Default is all origins
			c.allowedOriginsAll = true
		}
	} else {
		c.allowedOrigins = []string{}
		c.allowedWOrigins = []wildcard{}
		for _, origin := range options.AllowedOrigins {
			// Normalize
			origin = strings.ToLower(origin)
			if origin == "*" {
				// If "*" is present in the list, turn the whole list into a match all
				c.allowedOriginsAll = true
				c.allowedOrigins = nil
				c.allowedWOrigins = nil
				break
			} else if i := strings.IndexByte(origin, '*'); i >= 0 {
				// Split the origin in two: start and end string without the *
				w := wildcard{origin[0:i], origin[i+1:]}
				c.allowedWOrigins = append(c.allowedWOrigins, w)
			} else {
				c.allowedOrigins = append(c.allowedOrigins, origin)
			}
		}
	}

	// Allowed Headers
	if len(options.AllowedHeaders) == 0 {
		// Use sensible defaults
		c.allowedHeaders = []string{"Origin", "Accept", "Content-Type"}
	} else {
		// Origin is always appended as some browsers will always request for this header at preflight
		c.allowedHeaders = convert(append(options.AllowedHeaders, "Origin"), http.CanonicalHeaderKey)
		for _, h := range options.AllowedHeaders {
			if h == "*" {
				c.allowedHeadersAll = true
				c.allowedHeaders = nil
				break
			}
		}
	}

	// Allowed Methods
	if len(options.AllowedMethods) == 0 {
		// Default is spec's "simple" methods
		c.allowedMethods = []string{http.MethodGet, http.MethodPost, http.MethodHead}
	} else {
		c.allowedMethods = convert(options.AllowedMethods, strings.ToUpper)
	}

	return c
}

// Handler creates a new Cors handler with passed options.
func Handler(options Options) func(next http.Handler) http.Handler {
	c := New(options)
	return c.Handler
}

// AllowAll create a new Cors handler with permissive configuration allowing all
// origins with all standard methods with any header and credentials.
func AllowAll() *Cors {
	return New(Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})
}

// Handler apply the CORS specification on the request, and add relevant CORS headers
// as necessary.
func (c *Cors) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions && r.Header.Get("Access-Control-Request-Method") != "" {
			c.logf("Handler: Preflight request")
			c.handlePreflight(w, r)
			// Preflight requests are standalone and should stop the chain as some other
			// middleware may not handle OPTIONS requests correctly. One typical example
			// is authentication middleware ; OPTIONS requests won't carry authentication
			// headers (see #1)
			if c.optionPassthrough {
				next.ServeHTTP(w, r)
			} else {
				w.WriteHeader(http.StatusOK)
			}
		} else {
			c.logf("Handler: Actual request")
			c.handleActualRequest(w, r)
			next.ServeHTTP(w, r)
		}
	})
}

// handlePreflight handles pre-flight CORS requests
func (c *Cors) handlePreflight(w http.ResponseWriter, r *http.Request) {
	headers := w.Header()
	origin := r.Header.Get("Origin")

	if r.Method != http.MethodOptions {
		c.logf("Preflight aborted: %s!=OPTIONS", r.Method)
		return
	}
	// Always set Vary headers
	// see https://github.com/rs/cors/issues/10,
	//     https://github.com/rs/cors/commit/dbdca4d95feaa7511a46e6f1efb3b3aa505bc43f#commitcomment-12352001
	headers.Add("Vary", "Origin")
	headers.Add("Vary", "Access-Control-Request-Method")
	headers.Add("Vary", "Access-Control-Request-Headers")

	if origin == "" {
		c.logf("Preflight aborted: empty origin")
		return
	}
	if !c.isOriginAllowed(r, origin) {
		c.logf("Preflight aborted: origin '%s' not allowed", origin)
		return
	}

	reqMethod := r.Header.Get("Access-Control-Request-Method")
	if !c.isMethodAllowed(reqMethod) {
		c.logf("Preflight aborted: method '%s' not allowed", reqMethod)
		return
	}
	reqHeaders := parseHeaderList(r.Header.Get("Access-Control-Request-Headers"))
	if !c.areHeadersAllowed(reqHeaders) {
		c.logf("Preflight aborted: headers '%v' not allowed", reqHeaders)
		return
	}
	if c.allowedOriginsAll {
		headers.Set("Access-Control-Allow-Origin", "*")
	} else {
		headers.Set("Access-Control-Allow-Origin", origin)
	}
	// Spec says: Since the list of methods can be unbounded, simply returning the method indicated
	// by Access-Control-Request-Method (if supported) can be enough
	headers.Set("Access-Control-Allow-Methods", strings.ToUpper(reqMethod))
	if len(reqHeaders) > 0 {

		// Spec says: Since the list of headers can be unbounded, simply returning supported headers
		// from Access-Control-Request-Headers can be enough
		headers.Set("Access-Control-Allow-Headers", strings.Join(reqHeaders, ", "))
	}
	if c.allowCredentials {
		headers.Set("Access-Control-Allow-Credentials", "true")
	}
	if c.maxAge > 0 {
		headers.Set("Access-Control-Max-Age", strconv.Itoa(c.maxAge))
	}
	c.logf("Preflight response headers: %v", headers)
}

// handleActualRequest handles simple cross-origin requests, actual request or redirects
func (c *Cors) handleActualRequest(w http.ResponseWriter, r *http.Request) {
	headers := w.Header()
	origin := r.Header.Get("Origin")

	// Always set Vary, see https://github.com/rs/cors/issues/10
	headers.Add("Vary", "Origin")
	if origin == "" {
		c.logf("Actual request no headers added: missing origin")
		return
	}
	if !c.isOriginAllowed(r, origin) {
		c.logf("Actual request no headers added: origin '%s' not allowed", origin)
		return
	}

	// Note that spec does define a way to specifically disallow a simple method like GET or
	// POST. Access-Control-Allow-Methods is only used for pre-flight requests and the
	// spec doesn't instruct to check the allowed methods for simple cross-origin requests.
	// We think it's a nice feature to be able to have control on those methods though.
	if !c.isMethodAllowed(r.Method) {
		c.logf("Actual request no headers added: method '%s' not allowed", r.Method)

		return
	}
	if c.allowedOriginsAll {
		headers.Set("Access-Control-Allow-Origin", "*")
	} else {
		headers.Set("Access-Control-Allow-Origin", origin)
	}
	if len(c.exposedHeaders) > 0 {
		headers.Set("Access-Control-Expose-Headers", strings.Join(c.exposedHeaders, ", "))
	}
	if c.allowCredentials {
		headers.Set("Access-Control-Allow-Credentials", "true")
	}
	c.logf("Actual response added headers: %v", headers)
}

// convenience method. checks if a logger is set.
func (c *Cors) logf(format string, a ...interface{}) {
	if c.Log != nil {
		c.Log.Printf(format, a...)
	}
}

// isOriginAllowed checks if a given origin is allowed to perform cross-domain requests
// on the endpoint
func (c *Cors) isOriginAllowed(r *http.Request, origin string) bool {
	if c.allowOriginFunc != nil {
		return c.allowOriginFunc(r, origin)
	}
	if c.allowedOriginsAll {
		return true
	}
	origin = strings.ToLower(origin)
	for _, o := range c.allowedOrigins {
		if o == origin {
			return true
		}
	}
	for _, w := range c.allowedWOrigins {
		if w.match(origin) {
			return true
		}
	}
	return false
}

// isMethodAllowed checks if a given method can be used as part of a cross-domain request
// on the endpoint
func (c *Cors) isMethodAllowed(method string) bool {
	if len(c.allowedMethods) == 0 {
		// If no method allowed, always return false, even for preflight request
		return false
	}
	method = strings.ToUpper(method)
	if method == http.MethodOptions {
		// Always allow preflight requests
		return true
	}
	for _, m := range c.allowedMethods {
		if m == method {
			return true
		}
	}
	return false
}

// areHeadersAllowed checks if a given list of headers are allowed to used within
// a cross-domain request.
func (c *Cors) areHeadersAllowed(requestedHeaders []string) bool {
	if c.allowedHeadersAll || len(requestedHeaders) == 0 {
		return true
	}
	for _, header := range requestedHeaders {
		header = http.CanonicalHeaderKey(header)
		found := false
		for _, h := range c.allowedHeaders {
			if h == header {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
