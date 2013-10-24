<table width="900" border="1" cellspacing="0">
                 <tr>
                 	 <td>怪物 ID</td>                
                   <td>怪物名称</td>
				   <td>怪物血量</td>
                   <td>怪物攻击力</td>
                   <td>怪物防御力</td>
                   <td>怪物敏捷力</td>
				   <td>怪物抗性</td>
				   <td>怪物杀获经验</td>
                 </tr>
                 
                 {{range .List}}
                 <tr>
                   <td>{{.Id}}</td>           
                   <td>{{.Name}}</td>
                   <td>{{.HP}}</td>
                   <td>{{.Attack}}</td>
                   <td>{{.Defense}}</td>
				   <td>{{.Agility}}</td>
				   <td>{{.Resistance}}</td>
				   <td>{{.Exp}}</td>
                   <td><a href="/admin/monster/edit/{{.Id}}">修改</a>   <a href="/admin/monster/delete/{{.Id}}">删除</a></td>
                 </tr>
                 {{end}}                 
</table>