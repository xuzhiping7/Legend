 	Map ID:{{.Map.Id}}</br></br>		
	Map Number:{{.Map.Number}}</br>
    名字: {{.Map.Name}}</br>
    描述: {{.Map.MapDescript}}</br>
	需求等级:{{.Map.Level}}</br>
	
	</br>
	NPC列表：</br>
	</br>
	{{ range $i, $v :=  .Map.NPCs}}
		 {{$i}} : {{$v.NPCNumber}} </br>操作:<a href="/admin/map/delete_npc/{{$v.Id}}">删除</a> </br></br>
	{{end}}  
	 
	<form action="/admin/map/add_npc/{{.Map.Id}}" method="post">
	<input name="map_number" type="hidden" value="{{.Map.Number}}"></br>
	<input name="npc_number" type="text" value="添加NPC的ID"></br>
    <input type="submit" value="添加NPC" />
	</form>
	