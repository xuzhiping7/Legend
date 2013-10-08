<table width="900" border="1" cellspacing="0">
                 <tr>
                 	 <td>NPC ID</td>                
                   <td>NPC名称</td>
                   <td>类型</td>
                   <td>简单描述</td>
                   <td>详细描述</td>
                   <td>操作</td>
                 </tr>
                 
                 {{range .NPCList}}
                 <tr>
                   <td>{{.Id}}</td>           
                   <td>{{.Name}}</td>
                   <td>{{.Type}}</td>
                   <td>{{.DescSimple}}</td>
                   <td>{{.DescDetail}}</td>
                   <td><a href="/admin/npc/edit/{{.Id}}">修改</a>   <a href="/admin/npc/delete/{{.Id}}">删除</a></td>
                 </tr>
                 {{end}}                 
</table>