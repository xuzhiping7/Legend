package models

import (
	//"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//数据库对象
var OrmHandle orm.Ormer

func init() {
	//数据库对象初始化
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", "root:@/studygolang?charset=utf8", 30)
	orm.RegisterModel(new(Player), new(PlayerProp), new(LanguageTemplate), new(NPC),
		new(WechatMap), new(Monster), new(Prop), new(NPCConversation))

	OrmHandle = orm.NewOrm()

	//testtest := []PlayerProp{}
	//OrmHandle.QueryTable("wechat_player_prop").Filter("player_id", 1).All(&testtest)
	////OrmHandle.Where("id=1").FindAll(&testtest)
	////aa := GetPlayer("xuzhipingtest")
	//beego.Debug(testtest)
}
