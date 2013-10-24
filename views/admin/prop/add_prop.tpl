 <form action="/admin/prop/add" method="post"> 
	道具名称: <input name="name" type="text" value=""></br>
    道具描述：<input name="descript" type="text" value=""></br>
	道具出售价格: <input name="worth" type="text" value=""></br>
    道具购买价格：<input name="official_worth" type="text" value=""></br>
	道具种类: <input name="prop_type" type="text" value=""></br>
	<p>PropType_没有任何作用 = 0  </p>	
	<p>PropType_恢复生命值  = 1</p>
	<p>PropType_恢复行动力  = 2</p>
	<p>PropType_角色昵称更改 = 3</p>
	<p>PropType_角色确认洗点 = 4</p>
    道具使用的值：<input name="prop_value" type="text" value=""></br>	
    <input type="submit" value="添加" />
</form>