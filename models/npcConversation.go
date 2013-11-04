package models

import (
	"github.com/astaxie/beego"
)

/*
备注：这个NPCConversation设计以后可能会带来隐患，因为大部分都是用编号（NUMBER）的方式来辨识
	 而这个用了NPCId。以后需要注意是否会有异常的情况。
*/

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
		beego.Error("npc.go GetNPCConversations error:", err)
	}
	return &conversations
}
