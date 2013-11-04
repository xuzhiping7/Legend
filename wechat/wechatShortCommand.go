package wechat

import (
	"Legend/models"
	"fmt"
	//"github.com/astaxie/beego"
	"strconv"
)

//短命令处理函数
func playerShortCommandController(player *models.Player, content string) (s_ReturnContent string, isOnCommand bool) {
	isOnCommand = true
	s_ReturnContent = "test"

	//将字符串转换为数字
	number, err := strconv.ParseInt(content, 10, 32)

	if err == nil {
		commandNum := int(number)
		s_ReturnContent = ExecuteNumberCommand(player, commandNum)
	}

	return s_ReturnContent, isOnCommand
}

//执行地图中的数字命令
func ExecuteNumberCommand(player *models.Player, command int) (s string) {
	//beego.Trace("ExecuteNumberCommand")

	//玩家没有进入NPC短命令状态
	if player.LocationCommand == 0 {
		npcNumber := GetMapNPCNumberByShortCommand(player.Location, command)
		//beego.Trace(npcNumber)
		if npcNumber > 0 {
			//让玩家进入NPC短命令状态
			player.UpdateLocationCommand(command)
			s = ShowNPCDetail(npcNumber)
		} else {
			//取消NPC短命令状态
			player.UpdateLocationCommand(0)
			//如果找不到相应的命令，重新显示当前地图信息。
			s, _ = playerCommandController(player, commandPrefix[1])
		}
	} else {
		//执行NPC相应的短命令
		npcNumber := GetMapNPCNumberByShortCommand(player.Location, player.LocationCommand)
		if command == 1 {
			//NPC对话命令
			s = map_NPCs[npcNumber].Talk()
		} else if command == 0 {
			//返回命令
			//取消NPC短命令状态
			player.UpdateLocationCommand(0)
			//如果找不到相应的命令，重新显示当前地图信息。
			s, _ = playerCommandController(player, commandPrefix[1])
		} else {
			//没有这个命令，再次显示当前信息
			s = ShowNPCDetail(npcNumber)
		}

	}

	return s
}

//通过当前输入的短数字命令，获取需要执行NPC的编号，如果没有该命令，则返回0。
//因为记录的时候要避开0，所以储存的值实际上是数组引用的+1，这里要取得正确的值则要shortCommand - 1
func GetMapNPCNumberByShortCommand(mapNumber int, shortCommand int) (number int) {

	if map_MapData[mapNumber].NPCs != nil {
		for k, v := range *map_MapData[mapNumber].NPCs {
			if k == (shortCommand - 1) {
				return v.NPCNumber
			}
		}
	}
	return 0
}

//显示NPC的详细内容
func ShowNPCDetail(number int) (s string) {
	npc := map_NPCs[number]
	s = fmt.Sprintf("%s\n%s\n\n[1]交谈\n\n[0]返回", npc.Name, npc.DescDetail)
	return s
}
