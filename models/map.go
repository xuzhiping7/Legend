package models

import (
	"github.com/astaxie/beego"
)

type WechatMap struct {
	Id          int    `orm:"column(id);pk"`
	Number      int    `orm:"column(number)"`
	Name        string `orm:"column(name)"`
	MapDescript string `orm:"column(descript)"`
	Level       int    `orm:"column(level)"`
	NPCs        []int  `orm:"-"`
	Mosters     []int  `orm:"-"`
	MostersRate []int  `orm:"-"`
}

//为orm接口的表名
func (this *WechatMap) TableName() string {
	return "wechat_map"
}

//根据ID获取一个NPC
func GetOneMap(id int) (oneMap *WechatMap) {
	temp := WechatMap{}
	temp.Id = id
	err := OrmHandle.Read(&temp)
	if err != nil {
		beego.Error("map.go GetOneMap error:", err)
	}
	return &temp
}

//获取所有地图信息
func GetAllWechatMap() (allMap *[]WechatMap) {
	maps := []WechatMap{}
	_, err := OrmHandle.QueryTable("wechat_map").All(&maps)
	if err != nil {
		beego.Error("map.go GetAllWechatMap() :", err)
	}
	return &maps
}
