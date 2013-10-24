package models

import (
	"github.com/astaxie/beego"
)

//NPC对话
type NPCConversation struct {
	Id           int    `orm:"column(id);pk"`
	NPCId        int    `orm:"column(npc_id)"`
	Conversation string `orm:"column(conversation)"`
}

//为orm接口的表名
func (this *NPCConversation) TableName() string {
	return "wechat_npc_conversation"
}

//获取相应一个NPC的所有对白
func GetNPCConversations(id int) (list *[]NPCConversation) {
	conversations := []NPCConversation{}
	_, err := OrmHandle.QueryTable("wechat_npc_conversation").Filter("npc_id", id).All(&conversations)

	if err != nil {
		beego.Error("npc.go GetAllNPC error:", err)
	}
	return &conversations
}
