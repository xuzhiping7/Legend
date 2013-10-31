
 	NPC ID:{{.NPC.Id}}</br></br>
	
	NPC编号:{{.NPC.Number}}</br>
    名字: {{.NPC.Name}}</br>
    类型：{{.NPC.Type}}</br>
    简单描述: {{.NPC.DescSimple}}</br>
    详细: {{.NPC.DescDetail}}</br>
	
	</br>
	已有对白：</br>
	</br>
	{{ range $i, $v :=  .NPC.Conversations}}
		 {{$i}}  {{$v.Conversation}} </br>操作:<a href="/admin/npc/delete_conversation/{{$v.Id}}">删除</a> </br></br>
	{{end}}  
	 
	<form action="/admin/npc/add_conversation/{{.NPC.Id}}" method="post">
    <textarea name="conversation" cols="60" rows="5" >添加NPC对白的内容</textarea></br>
    <input type="submit" value="添加对白" />
	</form>
	