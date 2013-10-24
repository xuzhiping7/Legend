package models

import (
	"github.com/astaxie/beego"
)

//道具结构
type Prop struct {
	Id            int    `orm:"column(id);pk"`
	Name          string `orm:"column(name)"`
	Descript      string `orm:"column(descript)"`
	Worth         int    `orm:"column(worth)"`
	OfficialWorth int    `orm:"column(official_worth)"`
	PropType      int    `orm:"column(prop_type)"`
	PropValue     int    `orm:"column(prop_value)"`
}

//定义道具作用常量
const (
	PropType_没有任何作用 = 0
	PropType_恢复生命值  = 1
	PropType_恢复行动力  = 2
	PropType_角色昵称更改 = 3
	PropType_角色确认洗点 = 4
)

//为orm接口的表名
func (this *Prop) TableName() string {
	return "wechat_prop"
}

//获取指定ID的语言模板
func GetProp(id int) (prop *Prop) {
	temp := Prop{}
	temp.Id = id

	err := OrmHandle.Read(&temp)

	if err != nil {
		beego.Error("prop.go GetMonster error:", err)
	}
	return &temp
}

//获取所有地图信息
func GetAllProp() (all *[]Prop) {
	temps := []Prop{}
	_, err := OrmHandle.QueryTable("wechat_prop").All(&temps)
	if err != nil {
		beego.Error("monster.go GetAllWechatMap() :", err)
	}
	return &temps
}
