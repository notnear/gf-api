package common

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
)

/**
数据返回通用JSON数据结构
*/
type JsonResponse struct {
	Status  int         `json:"status"`         // 错误码((200:成功, 非200:失败))
	Message string      `json:"message"`        // 提示信息
	Data    interface{} `json:"data,omitempty"` // 返回数据(业务接口定义具体数据结构)
	Time    string      `json:"time"`           // 返回时间
}

/**
标准返回结果数据结构封装。
*/
func Json(r *ghttp.Request, err int, data interface{}) {
	message := GetErrorMessage(err)
	_ = r.Response.WriteJson(JsonResponse{
		Status:  err,
		Message: message,
		Data:    data,
		Time:    gtime.Now().Format("Y-m-d H:i:s.u"),
	})
}

/**
返回JSON数据并退出当前HTTP执行函数。
*/
func JsonExit(r *ghttp.Request, err int, data ...interface{}) {
	Json(r, err, data)
	r.Exit()
}

func JsonExitWithMessage(r *ghttp.Request, err int, message string, data ...interface{}) {
	responseData := interface{}(g.Map{})
	if len(data) > 0 {
		responseData = data[0]
	}
	_ = r.Response.WriteJson(JsonResponse{
		Status:  err,
		Message: message,
		Data:    responseData,
		Time:    gtime.Now().Format("Y-m-d H:i:s.u"),
	})
	r.Exit()
}
