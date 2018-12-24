package utils

import (
	"github.com/astaxie/beego/logs"
	"github.com/hugozhu/godingtalk"
	// "github.com/lengsh/godingding/log4go"
	// "log"
)

func init() {
}

// Model Struct
type Dingtalker struct {
	CorpSecret string
	CorpId     string
	AcToken    string
	ChatId     string
}

func NewDingtalker() *Dingtalker {
	// return &Dingtalker{ CORPSECRET, CORPID, ACTOKEN, CHATID}
	return &Dingtalker{ServerConfig.DcorpSecret, ServerConfig.DcorpId, ServerConfig.DacToken, ServerConfig.DchatId}
}

func (r *Dingtalker) SendChatTextMessage(msg string) {
	logs.Debug("corpSecret=", r.CorpSecret, "\ncorpId=", r.CorpId, "\nchatId=", r.ChatId)
	c := godingtalk.NewDingTalkClient(r.CorpId, r.CorpSecret)
	c.RefreshAccessToken()
	logs.Debug("AccessToken = ", c.AccessToken)
	err := c.SendTextMessage("YY", r.ChatId, msg)
	if err != nil {
		logs.Error(err)
	}
}

func (r *Dingtalker) SendChatLinkMessage(url string, img string, title string, text string) {
	logs.Debug("corpSecret=", r.CorpSecret, "\ncorpId=", r.CorpId, "\nchatId=", r.ChatId)
	c := godingtalk.NewDingTalkClient(r.CorpId, r.CorpSecret)
	c.RefreshAccessToken()
	logs.Debug("AccessToken = ", c.AccessToken)

	err := c.SendLinkMessage("YY", r.ChatId, img, url, title, text)
	//"http://47.105.107.171/query?do=report", "3大视频网站Comingsoon", "最新的优酷、腾讯、爱奇艺的热点电影近期上映及热度集中播报[Update]")
	if err != nil {
		logs.Error(err)
	}
}

func (r *Dingtalker) SendRobotTextMessage(msg string) {
	logs.Debug("corpSecret=", r.CorpSecret, "\ncorpId=", r.CorpId, "\nchatId=", r.ChatId)
	c := godingtalk.NewDingTalkClient(r.CorpId, r.CorpSecret)
	if c != nil {
		c.RefreshAccessToken()
		logs.Debug("AccessToken = ", c.AccessToken)
		err := c.SendRobotTextMessage(r.AcToken, msg)
		if err != nil {
			logs.Error(err)
		}
	}
}
