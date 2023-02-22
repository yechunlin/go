package jobs

import (
	"api/dao"
	"api/middleware/loggo"
	"api/util"
)

type Test struct{}

func (t Test) Run() {
	tmp := util.GetRandTwo(5, 10)
	nickname := util.GetRandStr(tmp)
	password := "151604"
	mobile := util.GetRandMobileNumber()
	token := util.MD5(nickname + password + mobile)
	user := dao.Users{
		Nickname: nickname,
		Mobile:   mobile,
		Token:    token,
		Password: util.MD5(password),
	}
	result := dao.CreateUser(&user)
	if result.Error != nil {
		loggo.WriteLogStr("插入失败" + nickname + ";" + mobile)
	}
}
