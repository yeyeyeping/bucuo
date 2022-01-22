package errormsg

const (
	Success                      = 200
	BadRequest                   = 400
	Unauthorized                 = 401
	InternalServerError          = 500
	Ok                           = 0
	ValidateError                = -1
	JwtError                     = -2
	ImperfectPersonalInformation = -3
	UnknowError                  = -4

	ResourceError = -5
)

var errmap = map[int]string{
	Success:                      "Success",
	BadRequest:                   "BadRequest",
	Unauthorized:                 "Unauthorized",
	Ok:                           "Success",
	ImperfectPersonalInformation: "请先补充个人信息",
	JwtError:                     "token有误",
	UnknowError:                  "出错了,工程师正在努力修复ing",
	ResourceError:                "资源归属错误！请刷新重试",
}

func GetMsg(code int, msg string) string {
	if msg != "" {
		return msg
	}
	if value, ok := errmap[code]; ok {
		return value
	} else {
		return msg
	}
}
