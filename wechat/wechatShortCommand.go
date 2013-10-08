package wechat

import (
	"Legend/models"
	"strconv"
)

func playerShortCommandController(player *models.Player, content string) (s_ReturnContent string, isOnCommand bool) {
	isOnCommand = true
	s_ReturnContent = "test"

	//将字符串转换为数字
	number, err := strconv.ParseInt(content, 10, 32)

	if err == nil {
		commandNum := int(number)
		s_ReturnContent = ExecuteNumberCommand(player.Location, commandNum)
	}

	return s_ReturnContent, isOnCommand
}

//执行地图中的数字命令
func ExecuteNumberCommand(mapIndex int, command int) (s string) {

	if 1 == 1 {
		s = "111"
	} else {
		s = "222"
	}
	return s
}
