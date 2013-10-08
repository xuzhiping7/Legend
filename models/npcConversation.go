package models

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
