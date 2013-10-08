package models

//玩家道具表
type PlayerProp struct {
	Id       int `orm:"column(id);pk"`
	PlayerId int `orm:"column(player_id)"`
	PropId   int `orm:"column(prop_id)"`
	PropNum  int `orm:"column(prop_num)"`
}

//为orm接口的表名
func (this *PlayerProp) TableName() string {
	return "wechat_player_prop"
}
