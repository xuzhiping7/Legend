<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=gb2312" />
<title>传说OL管理系统</title>
<link href="/static/css/style.css" rel="stylesheet" type="text/css" />
<script type="text/JavaScript"> 
var $=function(id) {
   return document.getElementById(id);
}

function show_menu(num){
for(i=0;i<100;i++){
	if($('li0'+i)){
	$('li0'+i).style.display='none';
	$('f0'+i).className='';
	 }
}
	  $('li0'+num).style.display='block';//触发以后信息块
	  $('f0'+num).className='left02down01_xia_li';//触发以后TAG样式
}

function show_menuB(numB){
	for(j=0;j<100;j++){
		 if(j!=numB){
			if($('Bli0'+j)){
		  $('Bli0'+j).style.display='none';
		  $('Bf0'+j).style.background='url(/static/img/01.gif)';
		}
		 }
	}
	if($('Bli0'+numB)){   
		if($('Bli0'+numB).style.display=='block'){
		  $('Bli0'+numB).style.display='none';
		 $('Bf0'+numB).style.background='url(/static/img/01.gif)';
		}else {
		  $('Bli0'+numB).style.display='block';
		  $('Bf0'+numB).style.background='url(/static/img/02.gif)';
		}
	}
}


var temp=0;
function show_menuC(){
		if (temp==0){
		 document.getElementById('LeftBox').style.display='none';
	  	 document.getElementById('RightBox').style.marginLeft='0';
		 document.getElementById('Mobile').style.background='url(/static/img/center.gif)';

		 temp=1;
		}else{
		document.getElementById('RightBox').style.marginLeft='222px';
	   	document.getElementById('LeftBox').style.display='block';
		document.getElementById('Mobile').style.background='url(/static/img/center0.gif)';

	   temp=0;
			}
	 }
	 
	 
</script>
</head>

<body>
<div class="header">
	<div class="header03"></div>
	<div class="header01"></div>
	<div class="header02"><a href="/admin">传说OL管理系统</a></div>
</div>
<div class="left" id="LeftBox">
	<div class="left01">
		<div class="left01_right"></div>
		<div class="left01_left"></div>
		<div class="left01_c">超级管理员：admin</div>
	</div>
	<div class="left02">
		<div class="left02top">
			<div class="left02top_right"></div>
			<div class="left02top_left"></div>
			<div class="left02top_c">用户信息管理</div>
		</div>
	  <div class="left02down">
			<div class="left02down01"><a  onclick="show_menuB(80)" href="javascript:;"><div id="Bf080" class="left02down01_img"></div>用户管理</a></div>
			<div class="left02down01_xia noneBox" id="Bli080">
				<ul>
					<li  id="f010"><a href="/admin/player">&middot;玩家列表</a></li>
					<li  id="f011"><a href="#">&middot;测试测试</a></li>
				</ul>
			</div>
			
		  <div class="left02down01"><a onclick="show_menuB(81)" href="javascript:;"><div id="Bf081" class="left02down01_img"></div>NPC管理</a></div>
			<div class="left02down01_xia noneBox" id="Bli081">
				<ul>
					<li  id="f012"><a href="/admin/npc">&middot;NPC列表</a></li>
					<li  id="f013"><a href="/admin/npc/add">&middot;新增NPC</a></li>
				</ul>
			</div>
			

			
			<div class="left02down01"><a onclick="show_menuB(82)" href="javascript:;"><div id="Bf082" class="left02down01_img"></div>地图管理</a></div>
			<div class="left02down01_xia noneBox" id="Bli082">
				<ul>
					<li  id="f012"><a href="/admin/map">&middot;地图列表</a></li>
					<li  id="f012"><a href="/admin/map/add">&middot;新增地图</a></li>
				</ul>
			</div>
			
			<div class="left02down01"><a onclick="show_menuB(83)" href="javascript:;"><div id="Bf083" class="left02down01_img"></div>语言对话模板</a></div>
			<div class="left02down01_xia noneBox" id="Bli083">
				<ul>
					<li  id="f012"><a href="/admin/textTemplate">&middot;对话模板列表</a></li>
					<li  id="f012"><a href="/admin/textTemplate/add">&middot;新建对话模板</a></li>
				</ul>
			</div>
			
			<div class="left02down01"><a onclick="show_menuB(84)" href="javascript:;"><div id="Bf084" class="left02down01_img"></div>怪物管理</a></div>
			<div class="left02down01_xia noneBox" id="Bli084">
				<ul>
					<li  id="f012"><a href="/admin/monster">&middot;怪物列表</a></li>
					<li  id="f013"><a href="/admin/monster/add">&middot;新增怪物</a></li>
				</ul>
			</div>
			
			
			<div class="left02down01"><a onclick="show_menuB(85)" href="javascript:;"><div id="Bf085" class="left02down01_img"></div>道具管理</a></div>
			<div class="left02down01_xia noneBox" id="Bli085">
				<ul>
					<li  id="f012"><a href="/admin/prop">&middot;道具列表</a></li>
					<li  id="f013"><a href="/admin/prop/add">&middot;新增道具</a></li>
				</ul>
			</div>
			
			
			
		</div>
	</div>
	
	<div class="left01">
		<div class="left03_right"></div>
		<div class="left01_left"></div>
		<div class="left03_c">安全退出</div>
	</div>
</div>
<div class="rrcc" id="RightBox">

	<div class="center" id="Mobile" onclick="show_menuC()"></div>
	<div class="right" id="li010">
		<div class="right01"><img src="/static/img/04.gif" /> 用户信息 &gt; <span>{{.Nav}}</span></div>
		{{.LayoutContent}}
	</div>
	
	
</div>
</body>
</html>
