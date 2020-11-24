package common

const (
	NotError     = 0
	SUCCESS      = 200
	ParamsError  = 300
	Forbid       = 403
	NoRoute      = 404
	SystemError  = 500
	AddDataError = iota + 10000
	AddUserFailed
	CheckDataFail
	UserHasFather
	UpdateUserFail
)

var codeMessage = map[int]string{
	NotError:       "无错误",
	SUCCESS:        "SUCCESS",
	ParamsError:    "参数错误",
	Forbid:         "请求被禁止",
	NoRoute:        "api not found",
	SystemError:    "system error",
	AddDataError:   "添加数据失败",
	AddUserFailed:  "添加用户失败",
	CheckDataFail:  "查询数据失败",
	UserHasFather:  "用户已有上级",
	UpdateUserFail: "更新用户失败",
}

func GetErrorMessage(code int) string {
	if vl, has := codeMessage[code]; has == true {
		return vl

	}
	return "not found value by key:" + string(code)
}
