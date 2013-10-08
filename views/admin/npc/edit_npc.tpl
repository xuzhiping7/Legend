 <form action="/admin/npc/edit/{{.NPC.Id}}" method="post">
 		ID:{{.NPC.Id}}</br>
    Name: <input name="name" type="text" value="{{.NPC.Name}}"></br>
    Type：<input name="type" type="text" value="{{.NPC.Type}}"></br>
    DescSimple: <input name="desc_simple" type="textarea" value="{{.NPC.DescSimple}}"></br>
    DescDetail: <textarea name="desc_detail" cols="60" rows="15" id="desc_detail">{{.NPC.DescDetail}}</textarea></br>
    <input type="submit" value="修改" />
</form>