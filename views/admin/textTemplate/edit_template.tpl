 <form action="/admin/textTemplate/edit/{{.Template.Id}}" method="post">
 		ID:{{.Template.Id}}</br>
    Number: <input name="number" type="text" value="{{.Template.Number}}"></br>
    内容: <textarea name="template" cols="60" rows="15" id="desc_detail">{{.Template.Template}}</textarea></br>
    <input type="submit" value="修改" />
</form>