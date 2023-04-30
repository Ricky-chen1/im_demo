package errno

//自定义错误
type Errno struct {
	Code    int
	Message string
}

//实现error接口
func (e *Errno) Error() string {
	return e.Message
}

const (
	Success     = 200
	ServerError = 500

	//用户相关
	UserNoExist       = 10001
	UserExist         = 10002
	UserCreateFail    = 10003
	UserLoginFail     = 10004
	PushRequestFail   = 10101
	AcceptRequestFail = 10102
	AddFriendFail     = 10103

	//群相关
	GroupCreateFail = 10201
	GroupJoinFail   = 10202

	//websocket相关
	WSConnectFali  = 20001
	WSConnectStop  = 20002
	WSReadSuccess  = 20003
	WSWriteSuccess = 20004

	//token鉴权相关
	NoToken        = 30001
	ParseTokenFail = 30002
	TokenExpired   = 30003

	//参数校验
	ParamsInvalid  = 40001
	ParamsBindFail = 40002
)

var CodeTag = map[int]string{
	Success: "success",

	UserNoExist:       "用户不存在",
	UserExist:         "用户已注册",
	UserCreateFail:    "用户注册失败",
	UserLoginFail:     "用户登陆失败",
	PushRequestFail:   "发送好友请求失败",
	AcceptRequestFail: "接受好友请求失败",
	AddFriendFail:     "添加好友失败",

	GroupCreateFail: "创建群聊失败",
	GroupJoinFail:   "加入群聊失败",

	WSConnectFali:  "websocket连接失败",
	WSConnectStop:  "websocket连接中断",
	WSReadSuccess:  "websocket写入信息成功",
	WSWriteSuccess: "websocket读取信息成功",

	NoToken:        "NO Token",
	ParseTokenFail: "token解析错误",
	TokenExpired:   "token已过期",

	ServerError:    "服务器内部错误",
	ParamsBindFail: "请求参数绑定失败",
	ParamsInvalid:  "请求参数不合法",
}
