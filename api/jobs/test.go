package jobs

import (
	"api/dao"
	"api/middleware/loggo"
	"api/util"
)

type Test struct{}

func (t Test) Run() {
	nickname := util.GetRandStr(5)
	password := "151604"
	mobile := util.GetRandMobileNumber()
	token := util.MD5(nickname + password + mobile)
	user := dao.Users{
		Nickname: nickname,
		Mobile:   mobile,
		Token:    token,
	}
	result := dao.CreateUser(&user)
	if result.Error != nil {
		loggo.WriteLogStr("插入失败" + nickname + ";" + mobile)
	}
}
