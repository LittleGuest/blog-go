package consts

const (
	StatusError                  = -1  // 服务器异常
	StatusOk                     = 0   // 正常
	StatusUserNotFound           = 001 // 用户不存在
	StatusAccountOrPasswordError = 002 // 账号或密码错误
	StatusEmailSendFail          = 003 // 邮件发送失败
)
