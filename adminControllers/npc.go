package adminControllers

import (
	"Legend/models"
	"github.com/astaxie/beego"
	"strconv"
)

type NPCShowController struct {
	beego.Controller
}

type NPCNewController struct {
	beego.Controller
}

type NPCEditController struct {
	beego.Controller
}

type NPCDeleteController struct {
	beego.Controller
}

func (this *NPCShowController) Get() {

	npcList := models.GetAllNPC()

	this.Data["Nav"] = "NPC列表"
	this.Data["NPCList"] = npcList

	this.Layout = "admin/admin.tpl"
	this.TplNames = "admin/npc/show_npc.tpl"
}

func (this *NPCNewController) Get() {
	this.Layout = "admin/admin.tpl"
	this.TplNames = "admin/npc/add_npc.tpl"
}

func (this *NPCNewController) Post() {

	postName := this.GetString("name")
	postDescSimple := this.GetString("desc_simple")
	postDescDetail := this.GetString("desc_detail")
	temp := this.Input().Get("type")
	postType, err := strconv.Atoi(temp)

	if err != nil {
		beego.Error("npc.go Post Err:", err)
	}

	npc := models.NPC{}
	npc.Name = postName
	npc.Type = postType
	npc.DescSimple = postDescSimple
	npc.DescDetail = postDescDetail

	_, err2 := models.OrmHandle.Insert(&npc)

	if err2 != nil {
		beego.Error("npc.go Post Err2:", err2)
	}

	this.Layout = "admin/admin.tpl"
	this.TplNames = "admin/npc/add_npc.tpl"
}

func (this *NPCEditController) Get() {

	id := this.Ctx.Input.Params(":id")
	intid, err := strconv.Atoi(id)

	if err != nil {
		beego.Error("npc.go Edit Get Err:", err)
	}

	npc := models.GetNPC(intid)

	this.Data["Nav"] = "NPC编辑"
	this.Data["NPC"] = npc

	this.Layout = "admin/admin.tpl"
	this.TplNames = "admin/npc/edit_npc.tpl"
}

func (this *NPCEditController) Post() {
	id := this.Ctx.Input.Params(":id")
	intid, err := strconv.Atoi(id)

	npc := models.GetNPC(intid)

	postName := this.GetString("name")
	postDescSimple := this.GetString("desc_simple")
	postDescDetail := this.GetString("desc_detail")
	temp := this.Input().Get("type")
	postType, err := strconv.Atoi(temp)

	if err != nil {
		beego.Error("npc.go Edit Err:", err)
	}

	npc.Name = postName
	npc.Type = postType
	npc.DescSimple = postDescSimple
	npc.DescDetail = postDescDetail

	_, err2 := models.OrmHandle.Update(npc)

	if err2 != nil {
		beego.Error("npc.go Edit Err2:", err2)
	}

	this.Data["Nav"] = "NPC编辑"
	this.Data["NPC"] = npc

	this.Layout = "admin/admin.tpl"
	this.TplNames = "admin/npc/edit_npc.tpl"
}

func (this *NPCDeleteController) Get() {

	id := this.Ctx.Input.Params(":id")
	intid, err := strconv.Atoi(id)

	if err != nil {
		beego.Error("npc.go Delete Get Err:", err)
	}

	npc := models.NPC{Id: intid}
	_, err2 := models.OrmHandle.Delete(&npc)

	if err2 != nil {
		beego.Error("npc.go Delete Get Err:", err)
	}
	this.Ctx.Redirect(302, "/admin/npc/list")
}
