package models

//道具结构
type Prop struct {
	Id            int
	Name          string
	Descript      string
	Worth         int
	OfficialWorth int
	PropType      int
	PropValue     int
}

//定义道具作用常量
const (
	PropType_没有任何作用 = 0
	PropType_恢复生命值  = 1
	PropType_恢复行动力  = 2
	PropType_角色昵称更改 = 3
	PropType_角色确认洗点 = 4
)
