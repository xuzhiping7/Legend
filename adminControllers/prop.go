package adminControllers

import (
	"Legend/models"
	"github.com/astaxie/beego"
	"strconv"
)

type PropController struct {
	beego.Controller
}

func (this *PropController) Get() {

	action := this.Ctx.Input.Params(":action")

	switch action {
	//列表
	case "":
		list := models.GetAllProp()

		this.Data["Nav"] = "道具列表"
		this.Data["List"] = list

		this.Layout = "admin/admin.tpl"
		this.TplNames = "admin/prop/list_prop.tpl"
	//添加
	case "add":
		this.Data["Nav"] = "添加道具模板"

		this.Layout = "admin/admin.tpl"
		this.TplNames = "admin/prop/add_prop.tpl"
	//编辑
	case "edit":
		id, err := strconv.Atoi(this.Ctx.Input.Params(":id"))

		if err != nil {
			beego.Error("prop.go Edit Get Err:", err)
		}
		prop := models.GetProp(id)
		this.Data["Nav"] = "模板编辑"
		this.Data["Prop"] = prop
		this.Layout = "admin/admin.tpl"
		this.TplNames = "admin/prop/edit_prop.tpl"

	//删除
	case "delete":
		id, err := strconv.Atoi(this.Ctx.Input.Params(":id"))

		if err != nil {
			beego.Error("prop.go Edit Get Err:", err)
		}

		temp := models.Prop{Id: id}
		_, err2 := models.OrmHandle.Delete(&temp)

		if err2 != nil {
			beego.Error("prop.go delete Get Err:", err)
		}
		this.Ctx.Redirect(302, "/admin/prop")
	}
}

func (this *PropController) Post() {

	action := this.Ctx.Input.Params(":action")

	switch action {
	//添加
	case "add":

		name := this.GetString("name")
		descript := this.GetString("descript")
		worth, err1 := strconv.Atoi(this.Input().Get("worth"))
		official_worth, err2 := strconv.Atoi(this.Input().Get("official_worth"))
		prop_type, err3 := strconv.Atoi(this.Input().Get("prop_type"))
		prop_value, err4 := strconv.Atoi(this.Input().Get("prop_value"))

		if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
			beego.Error("prop.go Post Err")
		}

		temp := models.Prop{}
		temp.Name = name
		temp.Descript = descript
		temp.Worth = worth
		temp.OfficialWorth = official_worth
		temp.PropType = prop_type
		temp.PropValue = prop_value

		_, err7 := models.OrmHandle.Insert(&temp)

		if err7 != nil {
			beego.Error("prop.go Post Err7:", err7)
		}

		this.Ctx.Redirect(302, "/admin/prop")

	//编辑
	case "edit":

		id, err0 := strconv.Atoi(this.Ctx.Input.Params(":id"))
		if err0 != nil {
			beego.Error("prop.go Edit Err:", err0)
		}

		name := this.GetString("name")
		descript := this.GetString("descript")
		worth, err1 := strconv.Atoi(this.Input().Get("worth"))
		official_worth, err2 := strconv.Atoi(this.Input().Get("official_worth"))
		prop_type, err3 := strconv.Atoi(this.Input().Get("prop_type"))
		prop_value, err4 := strconv.Atoi(this.Input().Get("prop_value"))

		if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
			beego.Error("prop.go Post Err")
		}

		temp := models.GetProp(id)
		temp.Name = name
		temp.Descript = descript
		temp.Worth = worth
		temp.OfficialWorth = official_worth
		temp.PropType = prop_type
		temp.PropValue = prop_value

		_, err7 := models.OrmHandle.Update(temp)

		if err7 != nil {
			beego.Error("prop.go Edit Err7:", err7)
		}
		this.Ctx.Redirect(302, "/admin/prop")

	}
}
