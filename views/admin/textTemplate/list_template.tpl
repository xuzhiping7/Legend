<table width="1000" border="1" cellspacing="0">
                 <tr>
                 	 <!--<td>ID</td> -->         
                   <td>编号</td>
                   <td>文字模板</td>
                   <td width="100">操作</td>
                 </tr>
                 {{range .TextList}}
                 <tr>
                   <!--<td>{{.Id}}</td>-->    
                   <td>{{.Number}}</td>
                   <td>{{.Template}}</td>
                   <td><a href="/admin/textTemplate/edit/{{.Id}}">修改</a>   <a href="/admin/textTemplate/delete/{{.Id}}">删除</a></td>
                 </tr>
                 {{end}}                 
</table>