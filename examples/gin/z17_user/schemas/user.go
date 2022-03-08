package schemas

// 用户结构体
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

// UserRegisterUsername 用户名注册结构体
type UserRegisterUsername struct {
	UserLoginUsername
	ResPassword string `json:"re_password"`
}

// UserLoginUsername 用户登录结构体
type UserLoginUsername struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// UserLoginEmail 用户登录结构体，根据邮箱登录
type UserLoginEmail struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserRegisterEmail 用户名注册结构体，根据邮箱注册
type UserRegisterEmail struct {
	UserLoginEmail
	ResPassword string `json:"re_password"`
}

// UserLoginPhone 用户登录结构体，根据手机号登录
type UserLoginPhone struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

// UserRegisterPhone 用户名注册结构体，根据手机号注册
type UserRegisterPhone struct {
	UserLoginPhone
	ResPassword string `json:"re_password"`
}

// UserLoginEmailCode 根据邮箱和验证码进行登录
type UserLoginEmailCode struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

// UserLoginPhoneCode 根据邮箱和验证码进行登录
type UserLoginPhoneCode struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}

type SendEmail struct {
	Email   []string `json:"email"`   // 邮箱列表
	Title   string   `json:"title"`   // 邮件标题
	Content string   `json:"content"` // 邮件内容
	IsHtml  bool     `json:"is_html"` // 是否为HTML邮件
	Cc      []string `json:"cc"`      // 抄送邮件列表
	Bcc     []string `json:"bcc"`     // 密送邮件列表
	Attach  uint64   `json:"attach"`  // 附件ID
}
