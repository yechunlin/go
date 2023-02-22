package controller

import (
	"api/conf"
	"api/dao"
	"api/model"
	"api/util"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
*
注册
*/
func Register(c *gin.Context) {
	nickname := c.DefaultPostForm("nickname", "")
	password := c.DefaultPostForm("password", "")
	mobile := c.DefaultPostForm("mobile", "")

	if nickname == "" || password == "" || mobile == "" {
		util.ReturnData(c, conf.API_ACCOUNT_ERROR, "账号有误", nil)
		return
	}

	token := util.MD5(nickname + password + mobile)
	user := dao.Users{
		Nickname: nickname,
		Mobile:   mobile,
		Token:    token,
	}
	result := dao.CreateUser(&user)
	if result.Error != nil {
		panic(result.Error)
	}
	// fmt.Println(result.RowsAffected)
	// fmt.Println(user.Id)
	// users := []dao.Users{
	// 	{Nickname: nickname, Mobile: mobile, Token: token},
	// 	{Nickname: nickname, Mobile: mobile, Token: token},
	// }
	// results := dao.CreateUserBatch(&users)
	// fmt.Println(results.RowsAffected)
	// for _, s := range users {
	// 	fmt.Println(s.Id)
	// }
	util.ReturnData(c, conf.API_SERVER_SUCCESS, "you are logged in", user)
}

/*
*

	查询用户信息
*/
func GetInfo(c *gin.Context) {
	id, err := strconv.ParseInt(c.DefaultPostForm("id", "0"), 10, 64)
	if err != nil {
		panic(err)
	}
	user := dao.Users{}
	result := dao.GetUserInfoById(&user, []string{"nickname"}, id)
	if result.Error != nil {
		panic(result.Error)
	}
	if result.RowsAffected <= 0 {
		util.ReturnData(c, conf.API_SERVER_SUCCESS, "无此用户", gin.H{})
		return
	}
	util.ReturnData(c, conf.API_SERVER_SUCCESS, "获取成功", user)
}

func GetUserList(c *gin.Context) {
	page := util.StrToInt8(c.DefaultPostForm("page", "1"))
	limit := util.StrToInt8(c.DefaultPostForm("limit", "5"))
	list := dao.ListPageResult{}
	result := dao.GetPage(&list, "", []string{}, "id desc", page, limit)
	fmt.Println(result.RowsAffected)
	util.ReturnData(c, conf.API_SERVER_SUCCESS, "获取成功", list)
}

func UpdateUser(c *gin.Context) {
	id := c.DefaultPostForm("id", "0")
	status := util.StrToInt8(c.DefaultPostForm("status", "1"))
	data := map[string]interface{}{
		"status": status,
	}
	model.DbInstance.Begin()
	//更新1
	result := dao.UpdateUser("id = "+id, data)
	if result.RowsAffected == 1 {
		model.DbInstance.Rollback()
		util.ReturnData(c, conf.API_SERVER_SUCCESS, "修改失败", gin.H{"affectRows": result.RowsAffected})
		return
	}

	//更新2
	result = dao.UpdateUser("id = 12", map[string]interface{}{
		"status": status,
	})
	if result.RowsAffected == 0 {
		model.DbInstance.Rollback()
		util.ReturnData(c, conf.API_SERVER_SUCCESS, "修改失败", gin.H{"affectRows": result.RowsAffected})
		return
	}
	model.DbInstance.Commit()
	util.ReturnData(c, conf.API_SERVER_SUCCESS, "修改成功", gin.H{"affectRows": result.RowsAffected})
}
