package constvar

//seq type
const (
	UserSeq = "user"
)

//uer status
const (
	UserNormal  = 100
	UserDeleted = 200
	UserFreezen = 300
)

//op record type
const (
	SignInOp = "SING_IN"
)

//register method
const (
	WeiXinRegister   = "WX"
	UsernameRegister = "USERNAME"
)

//login method
const (
	WeiXinLogin   = "WX"
	UsernameLogin = "USERNAME"
)

//login source
const (
	LoginFromWeb    = "Web"
	LoginFromWeChat = "WeChat"
)

//status
const (
	Normal  = 100
	Deleted = 200
)

//task state
const (
	TaskPaused    = 10
	TaskRunning   = 12
	TaskOverdue   = 14
	TaskCompleted = 15
)
