<table width="900" border="1" cellspacing="0">
                 <tr>
				   <td>道具ID</td>                
                   <td>道具名称</td>
				   <td>道具简介</td>
                   <td>道具售卖价格</td>
                   <td>道具购买价格</td>
                   <td>道具种类</td>
				   <td>道具使用的值</td>
                 </tr>
                 
                 {{range .List}}
                 <tr>
                   <td>{{.Id}}</td>           
                   <td>{{.Name}}</td>
                   <td>{{.Descript}}</td>
                   <td>{{.Worth}}</td>
                   <td>{{.OfficialWorth}}</td>
				   <td>{{.PropType}}</td>
				   <td>{{.PropValue}}</td>
                   <td><a href="/admin/prop/edit/{{.Id}}">修改</a>   <a href="/admin/prop/delete/{{.Id}}">删除</a></td>
                 </tr>
                 {{end}}                 
</table>