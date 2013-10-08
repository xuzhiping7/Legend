package controllers

import (
	"Legend/wechat"
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type WechatController struct {
	beego.Controller
}

//微信API接受结构
type textRecieveMessage struct {
	ToUserName   string
	FromUserName string
	CreateTime   string
	MsgType      string
	Content      string
	MsgId        string
}

//微信API返回结构
type textResponseMessage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   string
	MsgType      string
	Content      string
	FuncFlag     string
}

/*
	微信API通过GET请求过来的路由渲染只通过GET，得到结果后直接输出。
*/
func (this *WechatController) Get() {

	v := textRecieveMessage{}

	//本地开发
	bytes, _ := ioutil.ReadFile("conf/localtest.xml")

	//真实环境
	//bytes, _ := ioutil.ReadAll(req.Body)

	err := xml.Unmarshal(bytes, &v)
	if err != nil {
		beego.Error(err)
		return
	}

	responXML := textResponseMessage{}
	responXML.FromUserName = v.ToUserName
	responXML.ToUserName = v.FromUserName
	responXML.Content = wechat.WechatResponseHandle(v.FromUserName, v.Content)
	responXML.CreateTime = v.CreateTime
	responXML.MsgType = v.MsgType
	responXML.FuncFlag = "0"
	result, _ := xml.Marshal(responXML)

	fmt.Fprint(this.Ctx.ResponseWriter, string(result))
}

//beego貌似不能设置单独的控制器模板自动渲染，略过自动渲染的接口，这样才不会跟网页部分冲突。
func (this *WechatController) Render() error {
	return nil
}
