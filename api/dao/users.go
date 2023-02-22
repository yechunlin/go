package dao

import (
	"api/model"
	"math"

	"gorm.io/gorm"
)

type Users struct {
	Id         int64 `gorm:"primaryKey"`
	Nickname   string
	Mobile     string
	Token      string
	Password   string
	CreateTime string `gorm:"column:createTime;type:datetime"`
	Status     int    `gorm:"default:1"`
}

// 创建
func CreateUser(user *Users) *gorm.DB {
	return model.DbInstance.Select("Nickname", "Mobile", "Token", "Password").Create(&user)
}

// 批量创建
func CreateUserBatch(users *[]Users) *gorm.DB {
	return model.DbInstance.Select("Nickname", "Mobile", "Token", "Password").Create(&users)
}

// 查询 根据主键
func GetUserInfoById(user *Users, fields []string, id int64) *gorm.DB {
	return model.DbInstance.Select(fields).First(&user, id)
}

// 查询，单条
func GetOne(user *Users, where string, fields []string, order string) *gorm.DB {
	return model.DbInstance.Select(fields).Where(where).Order(order).Limit(1).Find(&user)
}

// 查询，获取全部 limit限制【0表示无限制】
func GetAll(user *[]Users, where string, fields []string, order string, limit int) *gorm.DB {
	result := model.DbInstance.Select(fields).Where(where).Order(order)
	if limit > 0 {
		result.Limit(limit)
	}
	return result.Find(&user)
}

// 分页查询
type ListPageResult struct {
	TotalSize   int64
	TotlePage   int64
	CurrentPage int64
	Data        []Users
}

func GetPage(list *ListPageResult, where string, fields []string, order string, page int, limit int) *gorm.DB {
	offect := (page - 1) * limit
	model.DbInstance.Model(&Users{}).Where(where).Count(&list.TotalSize)
	res := model.DbInstance.Select(fields).Where(where).Order(order).Offset(offect).Limit(limit).Find(&list.Data)
	tmp := float64(list.TotalSize) / float64(limit)
	list.TotlePage = int64(math.Ceil(tmp))
	list.CurrentPage = int64(page)
	return res
}

// 跟新
func UpdateUser(where string, data map[string]interface{}) *gorm.DB {
	return model.DbInstance.Model(&Users{}).Where(where).Updates(data)
}
