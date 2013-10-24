 <form action="/admin/monster/edit/{{.Monster.Id}}" method="post">
 	ID:{{.Monster.Id}}</br>	
	名称: <input name="name" type="text" value="{{.Monster.Name}}"></br>
    生命值：<input name="hp" type="text" value="{{.Monster.HP}}"></br>
	攻击力: <input name="attack" type="text" value="{{.Monster.Attack}}"></br>
    防御力：<input name="defense" type="text" value="{{.Monster.Defense}}"></br>
	敏捷: <input name="agility" type="text" value="{{.Monster.Agility}}"></br>
    抗性：<input name="resistance" type="text" value="{{.Monster.Resistance}}"></br>
	经验: <input name="exp" type="text" value="{{.Monster.Exp}}"></br>	
    <input type="submit" value="修改" />
</form>

