package adminControllers

import (
	"Legend/models"
	"github.com/astaxie/beego"
	"strconv"
)

type TextTemplateController struct {
	beego.Controller
}

func (this *TextTemplateController) Get() {

	action := this.Ctx.Input.Params(":action")

	switch action {
	//模板列表
	case "":
		textList := models.GetAllLanguageTemplate()

		this.Data["Nav"] = "对话模板列表"
		this.Data["TextList"] = textList

		this.Layout = "admin/admin.tpl"
		this.TplNames = "admin/textTemplate/list_template.tpl"
	//添加模板
	case "add":
		this.Data["Nav"] = "添加对话模板"

		this.Layout = "admin/admin.tpl"
		this.TplNames = "admin/textTemplate/add_template.tpl"
	//编辑模板
	case "edit":
		id := this.Ctx.Input.Params(":id")
		intid, err := strconv.Atoi(id)
		if err != nil {
			beego.Error("textTemplate.go Edit Get Err:", err)
		}
		template := models.GetLanguageTemplate(intid)
		this.Data["Nav"] = "模板编辑"
		this.Data["Template"] = template
		this.Layout = "admin/admin.tpl"
		this.TplNames = "admin/textTemplate/edit_template.tpl"
	//删除模板
	case "delete":
		id := this.Ctx.Input.Params(":id")
		intid, err := strconv.Atoi(id)
		if err != nil {
			beego.Error("textTemplate.go Edit Get Err:", err)
		}
		template := models.LanguageTemplate{Id: intid}
		_, err2 := models.OrmHandle.Delete(&template)

		if err2 != nil {
			beego.Error("textTemplate.go delete Get Err:", err)
		}
		this.Ctx.Redirect(302, "/admin/textTemplate")
	}
}

func (this *TextTemplateController) Post() {

	action := this.Ctx.Input.Params(":action")

	switch action {
	//添加模板
	case "add":

		template := this.GetString("template")

		temp := this.Input().Get("number")
		number, err := strconv.Atoi(temp)

		if err != nil {
			beego.Error("textTemplate.go Post Err:", err)
		}

		textTemplate := models.LanguageTemplate{}
		textTemplate.Number = number
		textTemplate.Template = template

		_, err2 := models.OrmHandle.Insert(&textTemplate)

		if err2 != nil {
			beego.Error("textTemplate.go Post Err2:", err2)
		}
		this.Ctx.Redirect(302, "/admin/textTemplate")

	//编辑模板
	case "edit":
		id := this.Ctx.Input.Params(":id")
		intid, err := strconv.Atoi(id)

		if err != nil {
			beego.Error("textTemplate.go Edit Err:", err)
		}

		template := models.GetLanguageTemplate(intid)

		tempNumber, err3 := this.GetInt("number")
		if err3 != nil {
			beego.Error("textTemplate.go Edit Err3:", err3)
		}
		template.Number = int(tempNumber)

		template.Template = this.GetString("template")

		_, err2 := models.OrmHandle.Update(template)

		if err2 != nil {
			beego.Error("textTemplate.go Edit Err2:", err2)
		}
		this.Ctx.Redirect(302, "/admin/textTemplate")

	}
}
