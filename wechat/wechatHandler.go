package wechat

import (
	"Legend/models"
	//"Legend/plugins"
	"github.com/astaxie/beego"
	//"strconv"
)

//在线玩家列表
var map_PlayerData map[string]*models.Player

//储存所有对话模板
var textTemplate map[int]string

//储存地图信息
var map_MapData map[int]*models.WechatMap

//储存所有NPC相应信息
var map_NPCs map[int]*models.NPC

//玩家等级对应经验信息
var LevelsExp *[30]int

func init() {
	//初始化玩家信息MAP
	map_PlayerData = make(map[string]*models.Player)

	//初始化所有对话模板
	textTemplate = make(map[int]string)

	lang := models.GetAllLanguageTemplate()

	for i := 0; i < len(lang); i++ {
		textTemplate[lang[i].Number] = lang[i].Template
	}

	//初始化所有地图信息
	map_MapData = models.GetAllMap()
	beego.Trace(map_MapData[3])

	//beego.Trace(textTemplate)

	//从csv读取数据（已完成，请勿打开，会重复插入很多数据的。）
	//testData := plugins.ReadCSV("conf/wechatTextTemplate.csv")
	//for i := 1; i < len(testData); i++ {
	//	textTemplate[testData[i][0]] = testData[i][1]

	//	number, _ := strconv.ParseInt(testData[i][0], 10, 32)
	//	lang := models.LanguageTemplate{Number: int(number), Template: testData[i][1]}
	//	id, err := models.OrmHandle.Insert(&lang)
	//	if err == nil {
	//		beego.Trace(lang, "  sucess:", id)
	//	} else {
	//		beego.Trace(lang, "  fail:", id)
	//	}

	//	//beego.Trace(lang)
	//}
	LevelsExp = &[30]int{10, 20, 40, 80, 100, 200, 300, 400, 500, 700,
		1000, 1200, 1400, 1800, 2000, 2500, 2700, 3000, 3500, 4000,
		5000, 5500, 6000, 6500, 7000, 7500, 8000, 8500, 9000, 10000}
}

func WechatResponseHandle(openid string, content string) (s_ReturnContent string) {

	/*
		假设不存在当前用户,让玩家进入注册流程
		存在当前用户，则读取当前用户的所有信息。
	*/
	player, ok := map_PlayerData[openid]

	if !ok {
		player = models.GetPlayer(openid)
		map_PlayerData[openid] = player
	}

	/*
		【流程一】：
		查看玩家是否有在事件节点上，如果有的话，返回相应字符串。若没有在事件节点，继续流程。
	*/
	s_ReturnContent, ok2 := playerFlagController(player, content)
	if ok2 {
		return s_ReturnContent
	}

	/*
		【流程二】：
		查看玩家输入的是否优先指令，是则执行，否则继续。
	*/

	s_ReturnContent, ok3 := playerCommandController(player, content)
	if ok3 {
		return s_ReturnContent
	}

	/*
		【流程三】：
		查看玩家输入的是否可用的快捷指令，是则执行，否则继续。
	*/

	s_ReturnContent, ok4 := playerShortCommandController(player, content)
	if ok4 {
		return s_ReturnContent
	}

	beego.Debug(s_ReturnContent)

	return ""
}
