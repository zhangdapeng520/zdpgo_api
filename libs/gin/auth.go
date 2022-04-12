package gin

import (
	"crypto/subtle"
	"encoding/base64"
	"net/http"
	"strconv"

	"github.com/zhangdapeng520/zdpgo_api/libs/gin/internal/bytesconv"
)

const AuthUserKey = "user"      // AuthUserKey 是基本身份验证中用户凭据的cookie名称。
type Accounts map[string]string // Accounts 定义用户名和密码的键值对
type authPair struct {          // 权限对
	value string
	user  string
}

type authPairs []authPair

// 查询权限
func (a authPairs) searchCredential(authValue string) (string, bool) {
	if authValue == "" {
		return "", false
	}
	for _, pair := range a {
		if subtle.ConstantTimeCompare([]byte(pair.value), []byte(authValue)) == 1 {
			return pair.user, true
		}
	}
	return "", false
}

// BasicAuthForRealm 返回HTTP Basic校验中间件。
// 它将map[string]字符串作为参数，其中键是用户名，值是密码，以及域的名称。
// 如果域为空，默认情况下将使用“需要授权”。
func BasicAuthForRealm(accounts Accounts, realm string) HandlerFunc {
	if realm == "" {
		realm = "Authorization Required"
	}
	realm = "Basic realm=" + strconv.Quote(realm)
	pairs := processAccounts(accounts)
	return func(c *Context) {
		// 查询允许通过的权限用户
		user, found := pairs.searchCredential(c.requestHeader("Authorization"))
		if !found {
			// 权益校验失败
			c.Header("WWW-Authenticate", realm)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// 如果找到了用户凭据，请将用户id设置为key AuthUserKey。在此上下文中，可以稍后使用 c.MustGet(gin.AuthUserKey).
		c.Set(AuthUserKey, user)
	}
}

// BasicAuth 返回HTTP权限校验中间件，只允许给定的账户列表通过请求
func BasicAuth(accounts Accounts) HandlerFunc {
	return BasicAuthForRealm(accounts, "")
}

// 处理账户列表
func processAccounts(accounts Accounts) authPairs {
	length := len(accounts)
	assert1(length > 0, "Empty list of authorized credentials")
	pairs := make(authPairs, 0, length)
	for user, password := range accounts {
		assert1(user != "", "User can not be empty")
		value := authorizationHeader(user, password)
		pairs = append(pairs, authPair{
			value: value,
			user:  user,
		})
	}
	return pairs
}

// 生成Basic权限校验请求头
func authorizationHeader(user, password string) string {
	base := user + ":" + password
	return "Basic " + base64.StdEncoding.EncodeToString(bytesconv.StringToBytes(base))
}
