package models

import (
	"math/rand"
)

import (
	"github.com/astaxie/beego"
)

const (
	NPCType_普通对话 = 0
	NPCType_武器店  = 1
	NPCType_防具店  = 2
)

//NPC通用接口
type WechatNPC interface {
	GetName() (s string)
	GetDescSimple() (s string)
	GetDescDetail() (s string)
	ShowNPC() (s string)
	Talk() (s string)
}

type NPC struct {
	Id         int    `orm:"column(id);pk"`
	Number     int    `orm:"column(number)"`
	Name       string `orm:"column(name)"`
	Type       int    `orm:"column(type)"`
	DescSimple string `orm:"column(desc_simple)"`
	DescDetail string `orm:"column(desc_detail)"`

	//NPC对话
	Conversations *[]NPCConversation `orm:"-"`
}

//为orm接口的表名
func (this *NPC) TableName() string {
	return "wechat_npc"
}

func (this *NPC) GetName() (s string) {
	return this.Name
}

func (this *NPC) GetDescSimple() (s string) {
	return this.DescSimple
}

func (this *NPC) GetDescDetail() (s string) {
	return this.DescDetail
}

func (this *NPC) Talk() (s string) {
	index := rand.Intn(len(*this.Conversations))
	return (*this.Conversations)[index].Conversation
}

func (this *NPC) ShowNPC() (s string) {
	return ""
}

//后台管理用到的
func GetNPC(id int) (npc *NPC) {
	temp := NPC{}
	temp.Id = id

	err := OrmHandle.Read(&temp)
	if err != nil {
		beego.Error("npc.go GetAllNPC error:", err)
	}

	temp.Conversations = GetNPCConversations(id)

	return &temp
}

func GetAllNPC() (npcList *[]NPC) {
	npcs := []NPC{}
	_, err := OrmHandle.QueryTable("wechat_npc").All(&npcs)

	if err != nil {
		beego.Error("npc.go GetAllNPC error:", err)
	}

	return &npcs
}
