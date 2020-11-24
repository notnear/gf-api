package controller

import (
	"gf-api/app/logic"
	"gf-api/common"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

// 更多校验规则请参考  https://goframe.org/util/gvalid/index
type AddRequest struct {
	Mobile   string `valid:"required|phone#电话不能为空|电话格式不正确" example:"13012345678" format:"string"`
	Nickname string `valid:"required|min-length:1#昵称不能为空|昵称格式不正确|昵称长度不正确" example:"张三" format:"string"`
	Avatar   string `valid:"required|min-length:1#头像不能为空|头像格式不正确|头像长度不正确" example:"https://www.baidu.com/img/flexible/logo/pc/result.png" format:"string"`
	Sex      int    `valid:"required|integer|min:0#性别不能为空|性别格式不正确|性别类型不正确" example:"1" format:"int"`
}

// @summary 添加用户
// @description 添加用户。
// @tags    用户
// @Accept  json
// @produce json
// @param   data formData AddRequest true "请求数据"
// @router  /user/add [POST]
// @success 200 {object} common.JsonResponse
// @Failure 400 {object} common.JsonResponse
func ActionAdd(r *ghttp.Request) {
	var requestData *AddRequest
	if err := r.GetStruct(&requestData); err != nil {
		common.JsonExitWithMessage(r, common.ParamsError, err.Error())
	}
	if e := gvalid.CheckStruct(requestData, nil); e != nil {
		common.JsonExitWithMessage(r, common.ParamsError, e.FirstString())
	}
	data, errCode := logic.AddUser(requestData.Mobile, requestData.Nickname, requestData.Avatar, requestData.Sex)
	if errCode != common.NotError {
		common.JsonExit(r, errCode)
	}
	common.Json(r, common.SUCCESS, g.Map{"userId": data})
}

type UserListRequest struct {
	Page     int `valid:"required|integer|min:1#页码不能为空|参数格式不正确|1类型不正确" example:"1" format:"int"`
	PageSize int `valid:"required|integer|min:1#size不能为空|参数格式不正确|2类型不正确" example:"10" format:"int"`
}

// @summary 用户列表
// @description 用户列表。
// @tags    用户
// @Accept  json
// @produce json
// @param   data formData UserListRequest true "请求数据"
// @router  /user/list [POST]
// @success 200 {object} common.JsonResponse
// @Failure 400 {object} common.JsonResponse
func ActionUserList(r *ghttp.Request) {
	var requestData *UserListRequest
	if err := r.GetStruct(&requestData); err != nil {
		common.JsonExitWithMessage(r, common.ParamsError, err.Error())
	}
	if e := gvalid.CheckStruct(requestData, nil); e != nil {
		common.JsonExitWithMessage(r, common.ParamsError, e.FirstString())
	}
	data, errCode := logic.FindUserList(requestData.Page, requestData.PageSize)
	if errCode != common.NotError {
		common.JsonExit(r, errCode)
	}
	common.Json(r, common.SUCCESS, data)
}
