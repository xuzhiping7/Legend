package wechat

import (
	"Legend/models"
	"fmt"
)

//定义枚举事件常量
const (
	flag_注册 = iota
	flag_注册完成
	flag_用户传入角色名申请更名操作
	flag_暂无
)

//查看玩家是否有在事件结点上，如果有的话，返回相应字符串，若不在的话，第二个参数会返回错误。
func playerFlagController(player *models.Player, content string) (s_ReturnContent string, isOnFlag bool) {
	isOnFlag = true
	s_ReturnContent = ""

	switch player.Flag {
	case flag_注册:
		//假设当前用户输入的不是注册
		if content == textTemplate[2] {
			s_ReturnContent = textTemplate[0]
			player.UpdateFlag(flag_注册完成)
		} else {
			s_ReturnContent = textTemplate[3]
		}

	case flag_注册完成:
		s_ReturnContent = textTemplate[4]
		player.UpdateFlag(flag_用户传入角色名申请更名操作)

	case flag_用户传入角色名申请更名操作:

		if player.UpdateNickName(content) {
			player.UpdateFlag(flag_暂无)
			s_ReturnContent = fmt.Sprintf(textTemplate[5], content)
		} else {
			s_ReturnContent = textTemplate[4]
		}

	case flag_暂无:
		isOnFlag = false
	}

	return s_ReturnContent, isOnFlag
}
