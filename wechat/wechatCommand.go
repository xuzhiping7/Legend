package wechat

import (
	"Legend/models"
	"fmt"
	//"github.com/astaxie/beego"
	"strconv"
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
		s_ReturnContent = ShowMapInfo(player.Location)

	//前往
	case strings.HasPrefix(content, commandPrefix[2]):

		str_AimMap := strings.Replace(content, commandPrefix[2], "", -1)

		//匹配玩家所在地
		mapId, b := GetMapNumberByMapName(str_AimMap)
		if b {
			if player.CommentPrefixStr != "" {
				player.CommentPrefixStr = ""
			}

			if map_MapData[mapId].Level <= player.Level {
				if player.UpdateLocation(mapId) {
					s_ReturnContent = fmt.Sprintf(textTemplate[800009], map_MapData[mapId].Name, map_MapData[mapId].MapDescript)
					player.RecordEvent.AddRecord(fmt.Sprintf(textTemplate[200001], map_MapData[mapId].Name))

				} else {
					s_ReturnContent = textTemplate[100]
				}
				//记录玩家前往事件
				player.RecordEvent.AddRecord(fmt.Sprintf(textTemplate[200001], map_MapData[mapId].Name))
			} else {
				s_ReturnContent = fmt.Sprintf(textTemplate[100028], map_MapData[mapId].Level)
			}
		} else {
			//如果没有匹配到地点，则输出当前玩家可以前往的地点，并且进入命令补全模式
			if player.CommentPrefixStr != "" {
				s_ReturnContent = textTemplate[7]
				player.CommentPrefixStr = ""
			} else {
				s_ReturnContent += textTemplate[10]
				for _, v := range map_MapData {
					if v.Level <= player.Level {
						s_ReturnContent += v.Name + "\n"
					}
				}
				player.CommentPrefixStr = commandPrefix[2]
			}
		}

	//修炼
	case strings.HasPrefix(content, commandPrefix[3]):
		////查看玩家当前地点是否适合修炼
		//b_CanMapPractice := CanMapPractice(player.Location, model.Func_修炼)
		//if !b_CanMapPractice {
		//	s_ReturnContent = textTemplate["800008"]
		//	break
		//}

		////查看是否有足够的行动力来执行改动作
		//temp_Str, b := PlayerCheckMobility(player, -5)
		//if b {
		//	//记录玩家修炼事件
		//	player.RecordEvent.AddRecord(fmt.Sprintf(textTemplate["200002"], Map_MapData[player.Location].Name))

		//	PlayerCheckStatus(player)
		//	s_ReturnContent = PlayerPratctice(player) + "\n\n" + temp_Str
		//} else {
		//	s_ReturnContent = temp_Str
		//}
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

//根据地图名字来获取地图信息
func GetMapNumberByMapName(name string) (index int, b bool) {
	b = false
	for _, v := range map_MapData {
		if name == v.Name {
			index = v.Number
			b = true
			break
		}
	}
	return index, b
}

//显示玩家所处当前地图的信息
func ShowMapInfo(number int) (s string) {
	s = fmt.Sprintf(textTemplate[200008], map_MapData[number].Name, map_MapData[number].MapDescript)
	//beego.Trace("ShowMapInfo")

	if map_MapData[number].NPCs != nil {
		for k, v := range *map_MapData[number].NPCs {
			//beego.Trace(k)
			s += "\n[" + strconv.Itoa(k+1) + "]" + map_NPCs[v.NPCNumber].GetName() + ":" + map_NPCs[v.NPCNumber].GetDescSimple()
		}
	}

	return s
}
