package Proto2

const (
	P2INIT          = iota
	Web_Login_Proto // Web_Login_Proto == 1 登陆
)

type Web_Login struct {
	Proto     int
	Proto2    int
	LoginName string // 用户名
	LoginPW   string // 密码
}
