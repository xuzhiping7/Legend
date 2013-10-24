package main

import (
	"Legend/adminControllers"
	"Legend/controllers"
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/orm"
)

func init() {
	//设置日志模块
	beego.SetLevel(beego.LevelTrace)
	beego.BeeLogger.SetLogger("file", `{"filename":"logs/logs.log"}`)
}

func main() {
	/*
		设置路由模块，并启动beego应用模块
	*/

	//网页模块
	beego.Router("/", &controllers.MainController{})

	//微信传说网游API模块
	beego.Router("/wechat:all", &controllers.WechatController{})
	beego.Router("/wechat", &controllers.WechatController{})

	//后台管理模块
	beego.Router("/admin", &adminControllers.AdminController{})
	beego.Router("/admin/", &adminControllers.AdminController{})

	beego.Router("/admin/player", &adminControllers.PlayerController{})

	beego.Router("/admin/npc", &adminControllers.NPCController{})
	beego.Router("/admin/npc/:action(add)", &adminControllers.NPCController{})
	beego.Router("/admin/npc/:action(detail|edit|delete|add_conversation|delete_conversation)/:id:int", &adminControllers.NPCController{})

	beego.Router("/admin/monster", &adminControllers.MonsterController{})
	beego.Router("/admin/monster/:action(add)", &adminControllers.MonsterController{})
	beego.Router("/admin/monster/:action(edit|delete)/:id:int", &adminControllers.MonsterController{})

	beego.Router("/admin/prop", &adminControllers.PropController{})
	beego.Router("/admin/prop/:action(add)", &adminControllers.PropController{})
	beego.Router("/admin/prop/:action(edit|delete)/:id:int", &adminControllers.PropController{})

	beego.Router("/admin/map", &adminControllers.MapController{})
	beego.Router("/admin/map/:action(add)", &adminControllers.MapController{})
	beego.Router("/admin/map/:action(edit|delete)/:id:int", &adminControllers.MapController{})

	beego.Router("/admin/textTemplate", &adminControllers.TextTemplateController{})
	beego.Router("/admin/textTemplate/:action(add)", &adminControllers.TextTemplateController{})
	beego.Router("/admin/textTemplate/:action(edit|delete)/:id:int", &adminControllers.TextTemplateController{})

	beego.Run()
}
