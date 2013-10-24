package models

import (
	"github.com/astaxie/beego"
)

type Monster struct {
	//野怪ID
	Id   int    `orm:"column(id);pk"`
	Name string `orm:"column(name)"`

	//野怪基本属性
	HP         int `orm:"column(hp)"`
	Attack     int `orm:"column(attack)"`
	Defense    int `orm:"column(defense)"`
	Agility    int `orm:"column(agility)"`
	Resistance int `orm:"column(resistance)"`

	//打败野怪可以获得的经验
	Exp int `orm:"column(exp)"`
}

//为orm接口的表名
func (this *Monster) TableName() string {
	return "wechat_monster"
}

//获取指定ID的语言模板
func GetMonster(id int) (monster *Monster) {
	temp := Monster{}
	temp.Id = id

	err := OrmHandle.Read(&temp)

	if err != nil {
		beego.Error("monster.go GetMonster error:", err)
	}
	return &temp
}

//获取所有地图信息
func GetAllMonster() (all *[]Monster) {
	temps := []Monster{}
	_, err := OrmHandle.QueryTable("wechat_monster").All(&temps)
	if err != nil {
		beego.Error("monster.go GetAllWechatMap() :", err)
	}
	return &temps
}
