<table width="900" border="1" cellspacing="0">
                 <tr>                
                   <td>玩家ID</td>
                   <td>昵称</td>
                   <td>等级</td>
                   <td>所在地</td>
                   <td>金钱</td>
                 </tr>
                 
                 {{range .PlayerList}}
                 <tr>                
                   <td>{{.Id}}</td>
                   <td>{{.NickName}}</td>
                   <td>{{.Level}}</td>
                   <td>{{.Location}}</td>
                   <td>{{.Coin}}</td>
                 </tr>
                 {{end}}                 
</table>