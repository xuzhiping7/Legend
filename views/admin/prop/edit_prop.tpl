 <form action="/admin/prop/edit/{{.Prop.Id}}" method="post">
 	ID:{{.Prop.Id}}</br>	
	道具名称: <input name="name" type="text" value="{{.Prop.Name}}"></br>
    道具描述：<input name="descript" type="text" value="{{.Prop.Descript}}"></br>
	道具出售价格: <input name="worth" type="text" value="{{.Prop.Worth}}"></br>
    道具购买价格：<input name="official_worth" type="text" value="{{.Prop.OfficialWorth}}"></br>
	道具种类: <input name="prop_type" type="text" value="{{.Prop.PropType}}"></br>
	<p>PropType_没有任何作用 = 0  </p>	
	<p>PropType_恢复生命值  = 1</p>
	<p>PropType_恢复行动力  = 2</p>
	<p>PropType_角色昵称更改 = 3</p>
	<p>PropType_角色确认洗点 = 4</p>
    道具使用的值：<input name="prop_value" type="text" value="{{.Prop.PropValue}}"></br>
    <input type="submit" value="修改" />
</form>

