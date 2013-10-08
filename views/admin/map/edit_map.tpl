 <form action="/admin/map/edit/{{.Map.Id}}" method="post">
 		ID:{{.Template.Id}}</br>
    Number: <input name="number" type="text" value="{{.Map.Number}}"></br>
    Level: <input name="level" type="text" value="{{.Map.Level}}"></br>
    Name: <input name="name" type="text" value="{{.Map.Name}}"></br>
    MapDescript: <textarea name="map_descript" cols="60" rows="15" >{{.Map.MapDescript}}</textarea></br>
    <input type="submit" value="修改" />
</form>