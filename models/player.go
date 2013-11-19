package models

import (
	"github.com/astaxie/beego"
	"time"
)

// 帖子信息
type Player struct {
	/*
		数据表中储存的数据
	*/
	Id               int    `orm:"column(id);pk"`
	OpenId           string `orm:"column(openid)"`
	NickName         string `orm:"column(nickname)"`
	UserName         string `orm:"column(username)"`
	Sex              int    `orm:"column(sex)"`
	Level            int    `orm:"column(level)"`
	Exp              int    `orm:"column(exp)"`
	Coin             int    `orm:"column(coin)"`
	Mobility         int    `orm:"column(mobility)"`
	Reputation       int    `orm:"column(reputation)"`
	Attack           int    `orm:"column(attack)"`
	Defense          int    `orm:"column(defense)"`
	Stamina          int    `orm:"column(stamina)"`
	Agility          int    `orm:"column(agility)"`
	Wisdom           int    `orm:"column(wisdom)"`
	NoDistribution   int    `orm:"column(no_distribution)"`
	Location         int    `orm:"column(location)"`
	LocationCommand  int    `orm:"column(location_command)"`
	LocationCommand2 int    `orm:"column(location_command2)"`
	Flag             int    `orm:"column(flag)"`

	/*
		扩展的动态数据
	*/
	//HP值
	Max_HP int `orm:"-"`
	Cur_HP int `orm:"-"`
	//当前行动力 `orm:"-"`
	Cur_Mobility int `orm:"-"`
	//当前抗性
	Cur_Resistance int `orm:"-"`
	//负重
	Max_Burden int `orm:"-"`
	Cur_Burden int `orm:"-"`
	//玩家道具
	Map_PlayerProp map[int]*PlayerProp `orm:"-"`

	/*
		辅助数据
	*/
	//当前命令前缀
	CommentPrefixStr string `orm:"-"`
	//当前前置事件提醒，例如被挑战
	ReplyPreifixStr string `orm:"-"`
	//玩家历史事件记录
	RecordEvent *PlayerRecord `orm:"-"`
	//玩家当前状态:0-无，1-休息
	Status int `orm:"-"`
	//玩家单独计时器
	Timer *time.Ticker `orm:"-"`
}

//为orm接口的表名
func (player *Player) TableName() string {
	return "wechat_base"
}

// 创建一个wechat玩家

func CreateWechatPlayer(openid string) (player *Player) {
	player = &Player{}

	player.OpenId = openid
	player.NickName = "EmptyNow"
	player.UserName = "EmptyNow"
	player.Exp = 0
	player.Mobility = 100
	player.Attack = 5
	player.Defense = 5
	player.Stamina = 5
	player.Agility = 5
	player.Wisdom = 5

	id, err := OrmHandle.Insert(player)
	player.Id = int(id)

	if err != nil {
		beego.Error("player.go CreateWechatPlayer(): insert player wrong !")
	}

	return player
}

//获取一个玩家信息
func GetPlayer(openid string) (player *Player) {
	//先获取临时的玩家
	tempPlayer := Player{OpenId: openid}
	err := OrmHandle.Read(&tempPlayer, "OpenId")
	if err != nil {
		beego.Error("player.go GetPlayer() ERR: ", err)
	}

	//若数据库中没有这个玩家，则新创建一个。
	if tempPlayer.Id == 0 {
		player = CreateWechatPlayer(openid)
	} else {
		player = &tempPlayer
	}

	//属性配点正确性验证
	if player.Attack+player.Agility+player.Wisdom+player.Stamina+player.Defense+player.NoDistribution > player.Level*3+25 {
		player.RedistributeAttribute()
		beego.Error("玩家数据异常，请查看玩家昵称：", player.NickName, "  ID:", player.Id)
	}

	//初始化动态信息
	player.InitPlayerProp()

	player.Cur_Mobility = player.Mobility
	player.Cur_HP = player.Stamina * 10
	player.Max_HP = player.Stamina * 10
	player.Cur_Resistance = player.Stamina + player.Defense*3 + player.Wisdom*3
	player.Max_Burden = player.Stamina * 5
	player.Cur_Burden = 0
	player.ReplyPreifixStr = ""
	player.RecordEvent = NewPlayerRecord()

	return player
}

//初始化玩家道具
func (player *Player) InitPlayerProp() {
	props := []PlayerProp{}
	_, err := OrmHandle.QueryTable("wechat_player_prop").Filter("player_id", 1).All(&props)

	if err != nil {
		beego.Error("player.go InitPlayerProp error:", err)
	} else {
		player.Map_PlayerProp = make(map[int]*PlayerProp)

		for _, prop := range props {
			player.Map_PlayerProp[prop.PropId] = &prop
		}
	}
}

//玩家重新分配点数，洗点
func (player *Player) RedistributeAttribute() {

	player.Attack = 5
	player.Defense = 5
	player.Agility = 5
	player.Stamina = 5
	player.Wisdom = 5

	player.NoDistribution = player.Level * 3

	_, err := OrmHandle.Update(player, "Attack", "Defense", "Agility", "Stamina", "Wisdom", "NoDistribution")

	if err != nil {
		beego.Error("player.go RedistributeAttribute(): update player wrong：", err)
	}

	player.Max_HP = player.Stamina * 10

	if player.Cur_HP > player.Stamina*10 {
		player.Cur_HP = player.Stamina * 10
	}

	player.Cur_Resistance = player.Stamina + player.Defense*3 + player.Wisdom*3
	player.Max_Burden = player.Stamina * 5
	player.Cur_Burden = 0
}

//玩家重命名
func (player *Player) UpdateNickName(name string) bool {

	runes := []rune(name)
	if len(runes) > 8 || len(runes) < 2 {
		return false
	}

	player.NickName = name
	_, err := OrmHandle.Update(player, "NickName")

	if err != nil {
		beego.Error("player.go Rename():", err)
		return false
	}

	return true
}

//更新玩家Flag
func (player *Player) UpdateFlag(flag int) bool {

	player.Flag = flag
	_, err := OrmHandle.Update(player, "Flag")

	if err != nil {
		beego.Error("player.go UpdateFlag():", err)
		return false
	}

	return true
}

//更新玩家位置
func (player *Player) UpdateLocation(i int) bool {
	//更新位置信息的同时需要更新短命令信息
	player.LocationCommand = 0
	player.Location = i
	_, err := OrmHandle.Update(player, "Location", "LocationCommand")
	if err != nil {
		beego.Error("player.go UpdateLocation():", err)
		return false
	}
	return true
}

//更新玩家位置当前的命令信息
func (player *Player) UpdateLocationCommand(i int) bool {
	player.LocationCommand = i
	_, err := OrmHandle.Update(player, "LocationCommand")
	if err != nil {
		beego.Error("player.go UpdateLocationCommand():", err)
		return false
	}
	return true
}

//更新玩家位置当前的命令信息（子命令）
func (player *Player) UpdateLocationCommand2(i int) bool {
	player.LocationCommand2 = i
	_, err := OrmHandle.Update(player, "LocationCommand2")
	if err != nil {
		beego.Error("player.go UpdateLocationCommand2():", err)
		return false
	}
	return true
}

//重置玩家位置当前的命令信息
func (player *Player) ResetLocationCommand() bool {
	if player.LocationCommand != 0 {
		player.UpdateLocationCommand(0)
	}
	if player.LocationCommand2 != 0 {
		player.UpdateLocationCommand2(0)
	}
	return true
}

//后台管理用到的
func GetAllPlayer() (playerList *[]Player) {
	players := []Player{}
	_, err := OrmHandle.QueryTable("wechat_base").All(&players)

	if err != nil {
		beego.Error("player.go GetAllPlayer error:", err)
	}
	return &players
}
