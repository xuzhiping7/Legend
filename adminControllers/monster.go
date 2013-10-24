package adminControllers

import (
	"Legend/models"
	"github.com/astaxie/beego"
	"strconv"
)

type MonsterController struct {
	beego.Controller
}

func (this *MonsterController) Get() {

	action := this.Ctx.Input.Params(":action")

	switch action {
	//列表
	case "":
		list := models.GetAllMonster()

		this.Data["Nav"] = "怪物列表"
		this.Data["List"] = list

		this.Layout = "admin/admin.tpl"
		this.TplNames = "admin/monster/list_monster.tpl"
	//添加
	case "add":
		this.Data["Nav"] = "添加怪物模板"

		this.Layout = "admin/admin.tpl"
		this.TplNames = "admin/monster/add_monster.tpl"
	//编辑
	case "edit":
		id, err := strconv.Atoi(this.Ctx.Input.Params(":id"))

		if err != nil {
			beego.Error("monster.go Edit Get Err:", err)
		}
		monster := models.GetMonster(id)
		this.Data["Nav"] = "模板编辑"
		this.Data["Monster"] = monster
		this.Layout = "admin/admin.tpl"
		this.TplNames = "admin/monster/edit_monster.tpl"

	//删除
	case "delete":
		id, err := strconv.Atoi(this.Ctx.Input.Params(":id"))

		if err != nil {
			beego.Error("monster.go Edit Get Err:", err)
		}

		monster := models.Monster{Id: id}
		_, err2 := models.OrmHandle.Delete(&monster)

		if err2 != nil {
			beego.Error("monster.go delete Get Err:", err)
		}
		this.Ctx.Redirect(302, "/admin/monster")
	}
}

func (this *MonsterController) Post() {

	action := this.Ctx.Input.Params(":action")

	switch action {
	//添加
	case "add":

		name := this.GetString("name")
		hp, err := strconv.Atoi(this.Input().Get("hp"))
		attack, err2 := strconv.Atoi(this.Input().Get("attack"))
		defense, err3 := strconv.Atoi(this.Input().Get("defense"))
		agility, err4 := strconv.Atoi(this.Input().Get("agility"))
		resistance, err5 := strconv.Atoi(this.Input().Get("resistance"))
		exp, err6 := strconv.Atoi(this.Input().Get("exp"))

		if err != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil || err6 != nil {
			beego.Error("monster.go Post Err")
		}

		monster := models.Monster{}
		monster.Name = name
		monster.HP = hp
		monster.Attack = attack
		monster.Defense = defense
		monster.Agility = agility
		monster.Resistance = resistance
		monster.Exp = exp

		_, err7 := models.OrmHandle.Insert(&monster)

		if err7 != nil {
			beego.Error("monster.go Post Err7:", err7)
		}

		this.Ctx.Redirect(302, "/admin/monster")

	//编辑
	case "edit":

		id, err0 := strconv.Atoi(this.Ctx.Input.Params(":id"))
		if err0 != nil {
			beego.Error("monster.go Edit Err:", err0)
		}

		name := this.GetString("name")
		hp, err := strconv.Atoi(this.Input().Get("hp"))
		attack, err2 := strconv.Atoi(this.Input().Get("attack"))
		defense, err3 := strconv.Atoi(this.Input().Get("defense"))
		agility, err4 := strconv.Atoi(this.Input().Get("agility"))
		resistance, err5 := strconv.Atoi(this.Input().Get("resistance"))
		exp, err6 := strconv.Atoi(this.Input().Get("exp"))

		if err != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil || err6 != nil {
			beego.Error("monster.go Post Err")
		}

		monster := models.GetMonster(id)
		monster.Name = name
		monster.HP = hp
		monster.Attack = attack
		monster.Defense = defense
		monster.Agility = agility
		monster.Resistance = resistance
		monster.Exp = exp

		_, err7 := models.OrmHandle.Update(monster)

		if err7 != nil {
			beego.Error("monster.go Edit Err7:", err7)
		}
		this.Ctx.Redirect(302, "/admin/monster")

	}
}
