package models

import (
	"github.com/astaxie/beego"
)

//NPC对话
type MapNPC struct {
	Id        int `orm:"column(id);pk"`
	MapNumber int `orm:"column(map_number)"`
	NPCNumber int `orm:"column(npc_number)"`
}

//为orm接口的表名
func (this *MapNPC) TableName() string {
	return "wechat_map_npc"
}

//获取相应一个地图的所有NPC编号
func GetMapNPC(id int) (list *[]MapNPC) {
	npcs := []MapNPC{}
	_, err := OrmHandle.QueryTable("wechat_map_npc").Filter("map_number", id).All(&npcs)

	if err != nil {
		beego.Error("mapNpc.go GetMapNPC error:", err)
	}
	return &npcs
}
