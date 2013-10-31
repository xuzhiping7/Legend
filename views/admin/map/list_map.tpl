<table width="1000" border="1" cellspacing="0">
                 <tr>
                 	 <!--<td>ID</td> -->  
                 	 <td>编号</td>
                   <td>地图名</td>
                   <td>准入等级</td>
                   <td>地图描述</td>
                   <td width="100">操作</td>
                 </tr>
                 {{range .MapList}}
                 <tr>
                   <!--<td>{{.Id}}</td>   -->
                   <td>{{.Number}}</td>
                   <td>{{.Name}}</td>
                   <td>{{.Level}}</td>
                   <td>{{.MapDescript}}</td>
                   <td><a href="/admin/map/edit/{{.Id}}">修改</a>   <a href="/admin/map/delete/{{.Id}}">删除</a> <a href="/admin/map/detail/{{.Id}}">详细</a></td>
                 </tr>
                 {{end}}                 
</table>