package logic

import (
	"gf-api/app/model/api_user"
	"gf-api/app/service"
	"gf-api/common"
	"github.com/gogf/gf/frame/g"
)

type UserLogic struct {
	service.UserService
}

func AddUser(mobile, nickname, avatar string, sex int) (userId int64, errCode int) {
	userLogic := UserLogic{}

	tx, err := g.DB().Begin()
	if err != nil {
		errCode = common.SystemError
		return
	}
	userData := api_user.Entity{
		Mobile:   mobile,
		Nickname: nickname,
		Avatar:   avatar,
		Sex:      sex,
	}

	userId, errCode = userLogic.AddUser(userData, tx)
	if errCode != common.NotError {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

func FindUserList(page, pageSize int) (data interface{}, errCode int) {
	userLogic := UserLogic{}
	where := g.Map{"status": 0}

	list, errCode := userLogic.FindUserList(where, page, pageSize)
	if errCode != common.NotError {
		return
	}

	total, errCode := userLogic.FindUserCount(where)
	if errCode != common.NotError {
		return
	}

	data = g.Map{
		"list":      list,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	}
	return
}
