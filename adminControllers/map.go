package adminControllers

import (
	"Legend/models"
	"github.com/astaxie/beego"
	"strconv"
)

type MapController struct {
	beego.Controller
}

func (this *MapController) Get() {

	action := this.Ctx.Input.Params(":action")

	switch action {
	//地图列表
	case "":
		maplist := models.GetAllWechatMap()

		this.Data["Nav"] = "地图列表"
		this.Data["MapList"] = maplist

		this.Layout = "admin/admin.tpl"
		this.TplNames = "admin/map/list_map.tpl"
	//添加地图
	case "add":
		this.Data["Nav"] = "添加地图"

		this.Layout = "admin/admin.tpl"
		this.TplNames = "admin/map/add_map.tpl"
	//编辑地图
	case "edit":
		id := this.Ctx.Input.Params(":id")
		intid, err := strconv.Atoi(id)
		if err != nil {
			beego.Error("map.go Edit Get Err:", err)
		}
		temp := models.GetOneMap(intid)

		this.Data["Nav"] = "地图编辑"
		this.Data["Map"] = temp
		this.Layout = "admin/admin.tpl"
		this.TplNames = "admin/map/edit_map.tpl"
	//删除地图
	case "delete":
		id := this.Ctx.Input.Params(":id")
		intid, err := strconv.Atoi(id)
		if err != nil {
			beego.Error("map.go Edit Get Err:", err)
		}
		tempMap := models.WechatMap{Id: intid}
		_, err2 := models.OrmHandle.Delete(&tempMap)

		if err2 != nil {
			beego.Error("map.go delete Get Err:", err)
		}
		this.Ctx.Redirect(302, "/admin/map")
	}
}

func (this *MapController) Post() {

	action := this.Ctx.Input.Params(":action")

	switch action {
	//添加地图
	case "add":

		name := this.GetString("name")
		descript := this.GetString("map_descript")

		temp := this.Input().Get("level")
		level, err := strconv.Atoi(temp)
		if err != nil {
			beego.Error("map.go Post Err:", err)
		}

		temp2 := this.Input().Get("number")
		number, err2 := strconv.Atoi(temp2)
		if err2 != nil {
			beego.Error("map.go Post Err2:", err2)
		}

		tempMap := models.WechatMap{}
		tempMap.Name = name
		tempMap.MapDescript = descript

		tempMap.Level = level
		tempMap.Number = number

		_, err3 := models.OrmHandle.Insert(&tempMap)

		if err3 != nil {
			beego.Error("map.go Post Err3:", err3)
		}
		this.Ctx.Redirect(302, "/admin/map")

	//编辑地图
	case "edit":

		id := this.Ctx.Input.Params(":id")
		intid, err := strconv.Atoi(id)

		if err != nil {
			beego.Error("map.go Edit Err:", err)
		}

		name := this.GetString("name")
		descript := this.GetString("map_descript")

		tempLevel, err2 := this.GetInt("level")
		if err2 != nil {
			beego.Error("map.go Edit Err2:", err2)
		}
		level := int(tempLevel)

		tempNumber, err3 := this.GetInt("number")
		if err3 != nil {
			beego.Error("map.go Edit Err2:", err3)
		}
		number := int(tempNumber)

		temp := models.GetOneMap(intid)
		temp.Number = number
		temp.Name = name
		temp.MapDescript = descript
		temp.Level = level

		_, err4 := models.OrmHandle.Update(temp)
		if err4 != nil {
			beego.Error("map.go Edit Err4:", err4)
		}

		this.Ctx.Redirect(302, "/admin/map")
	}
}
