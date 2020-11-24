package dao

import (
	"gf-api/app/model/api_user"
	"gf-api/common"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

type UserDao struct{}

func (u *UserDao) AddUser(data api_user.Entity, tx *gdb.TX) (userId int64, errCode int) {
	result, er := api_user.Model.TX(tx).Insert(data)
	if er != nil {
		errCode = common.AddDataError
		return
	}
	userId, _ = result.LastInsertId()
	return
}

func (u *UserDao) FindUserList(data g.Map, page, pageSize int) (source []*api_user.Entity, errCode int) {
	office := (page - 1) * pageSize
	source, er := api_user.Model.Page(page, office).FindAll(data)
	if er != nil {
		errCode = common.CheckDataFail
		return
	}
	return
}

func (u *UserDao) FindUserCount(data g.Map) (source int, errCode int) {
	source, er := api_user.Model.Count(data)
	if er != nil {
		errCode = common.CheckDataFail
		return
	}
	return
}
