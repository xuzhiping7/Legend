package adminControllers

import (
	"Legend/models"
	"github.com/astaxie/beego"
	"strconv"
)

type NPCController struct {
	beego.Controller
}

func (this *NPCController) Get() {
	action := this.Ctx.Input.Params(":action")

	switch action {
	//列表
	case "":
		npcList := models.GetAllNPC()

		this.Data["Nav"] = "NPC列表"
		this.Data["NPCList"] = npcList

		this.Layout = "admin/admin.tpl"
		this.TplNames = "admin/npc/list_npc.tpl"
	//添加
	case "add":
		this.Data["Nav"] = "添加NPC模板"

		this.Layout = "admin/admin.tpl"
		this.TplNames = "admin/npc/add_npc.tpl"
	//编辑
	case "edit":
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

	case "detail":
		id := this.Ctx.Input.Params(":id")
		intid, err := strconv.Atoi(id)

		if err != nil {
			beego.Error("npc.go Edit Get Err:", err)
		}

		npc := models.GetNPC(intid)

		this.Data["Nav"] = "NPC详细信息"
		this.Data["NPC"] = npc

		this.Layout = "admin/admin.tpl"
		this.TplNames = "admin/npc/detail_npc.tpl"

	//删除
	case "delete":
		id := this.Ctx.Input.Params(":id")
		intid, err := strconv.Atoi(id)

		if err != nil {
			beego.Error("npc.go Delete Get Err:", err)
		}

		npc := models.NPC{Id: intid}
		_, err2 := models.OrmHandle.Delete(&npc)

		if err2 != nil {
			beego.Error("npc.go Delete Get Err:", err2)
		}
		this.Ctx.Redirect(302, "/admin/npc")

	case "delete_conversation":
		id := this.Ctx.Input.Params(":id")
		intid, err := strconv.Atoi(id)

		if err != nil {
			beego.Error("npc.go delete_conversation Get Err:", err)
		}

		temp := models.NPCConversation{Id: intid}
		_, err2 := models.OrmHandle.Delete(&temp)

		if err2 != nil {
			beego.Error("npc.go delete_conversation Get Err:", err2)
		}
		this.Ctx.Redirect(302, "/admin/npc")
	}
}

func (this *NPCController) Post() {

	action := this.Ctx.Input.Params(":action")

	switch action {
	//添加
	case "add":

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

		this.Ctx.Redirect(302, "/admin/npc")

	//编辑
	case "edit":
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

		this.Ctx.Redirect(302, "/admin/npc")

	case "add_conversation":

		id, err := strconv.Atoi(this.Ctx.Input.Params(":id"))
		if err != nil {
			beego.Error("npc.go post action add_conversation Err:", err)
		}

		conversation := this.GetString("conversation")

		temp := models.NPCConversation{}
		temp.NPCId = id
		temp.Conversation = conversation

		_, err2 := models.OrmHandle.Insert(&temp)

		if err2 != nil {
			beego.Error("npc.go post action add_conversation Err2:", err2)
		}

		this.Ctx.Redirect(302, "/admin/npc/detail/"+strconv.Itoa(id))
	}
}
