package wechat

import (
	"Legend/models"
	"fmt"
	"strings"
)

//储存所有命令前缀
var commandPrefix map[int]string

func init() {

	commandPrefix = make(map[int]string)
	commandPrefix[0] = "我"
	commandPrefix[1] = "当前"
	commandPrefix[2] = "前往"
	commandPrefix[3] = "修炼"
	commandPrefix[4] = "状态"
	commandPrefix[5] = "搜寻"
	commandPrefix[6] = "帮助"
	commandPrefix[7] = "传说"
	commandPrefix[8] = "装备"
	commandPrefix[9] = "道具"
	commandPrefix[10] = "使用"
	commandPrefix[11] = "提升"
	commandPrefix[12] = "提升攻击"
	commandPrefix[13] = "提升体力"
	commandPrefix[14] = "提升防御"
	commandPrefix[15] = "提升敏捷"
	commandPrefix[16] = "提升智慧"
	commandPrefix[17] = "附近"
	commandPrefix[18] = "休息"
	commandPrefix[19] = "买"
	commandPrefix[20] = "卖"
	commandPrefix[21] = "挑战"
	//commandPrefix[22] = "确认洗点"
	commandPrefix[23] = "事件"
}

//查看玩家是否有在事件结点上，如果有的话，返回相应字符串，若不在的话，第二个参数会返回错误。
func playerCommandController(player *models.Player, content string) (s_ReturnContent string, isOnCommand bool) {
	isOnCommand = true
	s_ReturnContent = ""

	//如果有命令补全，则补全命令
	if player.CommentPrefixStr != "" {
		content = player.CommentPrefixStr + content
	}

	switch {
	//我
	case strings.HasPrefix(content, commandPrefix[0]):
		s_ReturnContent = fmt.Sprintf(textTemplate[6], player.NickName, map_MapData[player.Location].Name, player.Level, player.Exp, LevelsExp[player.Level], player.Cur_Mobility, player.Mobility, player.Reputation, player.Coin, player.Cur_HP, player.Max_HP, player.Cur_Burden, player.Max_Burden, player.Cur_Resistance, player.Attack, player.Defense, player.Stamina, player.Agility, player.Wisdom, player.NoDistribution)

	//当前
	case strings.HasPrefix(content, commandPrefix[1]):

	//前往
	case strings.HasPrefix(content, commandPrefix[2]):

	//修炼
	case strings.HasPrefix(content, commandPrefix[3]):

	//帮助
	case strings.HasPrefix(content, commandPrefix[6]):
		s_ReturnContent = textTemplate[11]

	//传说
	case strings.HasPrefix(content, commandPrefix[7]):

	//道具
	case strings.HasPrefix(content, commandPrefix[9]):

	//使用
	case strings.HasPrefix(content, commandPrefix[10]):

	//提升
	case content == commandPrefix[11]:

	//提升攻击
	case strings.HasPrefix(content, commandPrefix[12]):

	//提升体力
	case strings.HasPrefix(content, commandPrefix[13]):

	//提升防御
	case strings.HasPrefix(content, commandPrefix[14]):

	//提升敏捷
	case strings.HasPrefix(content, commandPrefix[15]):

	//提升智慧
	case strings.HasPrefix(content, commandPrefix[16]):

	//显示当前所在地图所有玩家
	case strings.HasPrefix(content, commandPrefix[17]):

	//进入休息状态
	case strings.HasPrefix(content, commandPrefix[18]):

	case strings.HasPrefix(content, commandPrefix[19]):

	//卖东西
	case strings.HasPrefix(content, commandPrefix[20]):

	//挑战玩家
	case strings.HasPrefix(content, commandPrefix[21]):

	//我确认洗点
	/*
		case strings.HasPrefix(content, commandPrefix[22]):
			//所有属性重置为5点
			s_ReturnContent = PlayerRedistributeAttribute(player)
	*/

	//事件
	case strings.HasPrefix(content, commandPrefix[23]):

	default:
		isOnCommand = false
	}

	//if player.ReplyPreifixStr != "" {
	//	s_ReturnContent = AddReplyPrefix(s_ReturnContent, player.ReplyPreifixStr)
	//	PlayerClearReplyPrefix(player)
	//}

	return s_ReturnContent, isOnCommand
}

//显示玩家所处当前地图的信息
func ShowMapInfo(player *models.Player) (s string) {

	s = fmt.Sprintf(textTemplate[200008], map_MapData[player.Location].Name, map_MapData[player.Location].MapDescript)

	if len(map_MapData[player.Location].NPCs) > 0 {
		for _, v := range map_MapData[player.Location].NPCs {
			s += "\n[1]" + map_NPCs[v].GetName() + ":" + map_NPCs[v].GetDescSimple()
		}
	}

	return s
}