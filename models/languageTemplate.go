package models

import (
	"github.com/astaxie/beego"
)

//玩家道具表
type LanguageTemplate struct {
	Id       int    `orm:"column(id);pk"`
	Number   int    `orm:"column(number)"`
	Template string `orm:"column(template)"`
}

//为orm接口的表名
func (this *LanguageTemplate) TableName() string {
	return "wechat_language_template"
}

//获取指定ID的语言模板
func GetLanguageTemplate(id int) (text *LanguageTemplate) {
	temp := LanguageTemplate{}
	temp.Id = id

	err := OrmHandle.Read(&temp)

	if err != nil {
		beego.Error("languageTempalte.go GetLanguageTemplate error:", err)
	}
	return &temp
}

//获取所有语言模板
func GetAllLanguageTemplate() (all []LanguageTemplate) {

	_, err := OrmHandle.QueryTable("wechat_language_template").OrderBy("number").All(&all)

	if err != nil {
		beego.Error("languageTemplate.go GetAllLanguageTemplate Error:", err)
	}

	return all
}
