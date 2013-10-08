package adminControllers

import (
	"Legend/models"
	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
}

func (this *AdminController) Get() {
	//this.Data["Nav"] = "玩家列表"

	this.TplNames = "admin/admin.tpl"

	//this.TplNames = "admin/player/playerList.tpl"
}

type PlayerController struct {
	beego.Controller
}

func (this *PlayerController) Get() {

	playerList := models.GetAllPlayer()

	this.Data["Nav"] = "玩家列表"
	this.Data["PlayerList"] = playerList

	this.Layout = "admin/admin.tpl"

	this.TplNames = "admin/player/playerList.tpl"
}
