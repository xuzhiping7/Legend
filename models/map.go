package models

import (
	"github.com/astaxie/beego"
)

type WechatMap struct {
	Id          int       `orm:"column(id);pk"`
	Number      int       `orm:"column(number)"`
	Name        string    `orm:"column(name)"`
	MapDescript string    `orm:"column(descript)"`
	Level       int       `orm:"column(level)"`
	NPCs        *[]MapNPC `orm:"-"`
	Mosters     []int     `orm:"-"`
	MostersRate []int     `orm:"-"`
}

//为orm接口的表名
func (this *WechatMap) TableName() string {
	return "wechat_map"
}

//根据ID获取一张地图信息
func GetOneMap(id int) (oneMap *WechatMap) {
	temp := WechatMap{}
	temp.Id = id
	err := OrmHandle.Read(&temp)
	if err != nil {
		beego.Error("map.go GetOneMap error:", err)
	}

	temp.NPCs = GetMapNPC(temp.Number)

	return &temp
}

//获取所有地图信息
func GetAllWechatMap() (allMap *[]WechatMap) {
	maps := []WechatMap{}
	_, err := OrmHandle.QueryTable("wechat_map").All(&maps)
	if err != nil {
		beego.Error("map.go GetAllWechatMap() :", err)
	}

	for i := 0; i < len(maps); i++ {
		maps[i].NPCs = GetMapNPC(maps[i].Number)
	}

	return &maps
}

//获取所有地图信息，以map[int]*models.WechatMap结构返回
func GetAllMap() (allMap map[int]*WechatMap) {

	//这个指针取值方式可能会有效率上的一点点缺失，但没到好的解决方法，除非再写个函数。
	maps := *GetAllWechatMap()

	allMap = make(map[int]*WechatMap)

	for i := 0; i < len(maps); i++ {
		allMap[maps[i].Number] = &maps[i]
	}

	return allMap
}
