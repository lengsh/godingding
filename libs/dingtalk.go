package libs

import (
	"github.com/hugozhu/godingtalk"
	"github.com/lengsh/godingding/log4go"
	// "log"
)

const CORPSECRET string = "2uK2a27AWgkfkVAxd9IdwqG9SO7D01LhWnCgDEYhxff6uGj924NEdrboCivL_Gry"
const CORPID string = "ding5b26ca68f242cff035c2f4657eb6378f"
const ACTOKEN string = "167ff2dd88c92f9267628960d78cd23fd0300d7f691d24631355f170df2a74cb"
const CHATID string = "chat8890dbc9d98595c5a1031fe99d8c585e"

func init() {

}

// Model Struct
type Dingtalker struct {
	CorpSecret string
	CorpId     string
	AcToken    string
	ChatId     string
	Loger      *log4go.G4Log
}

func NewDingtalker() *Dingtalker {
	return &Dingtalker{CORPSECRET, CORPID, ACTOKEN, CHATID, nil}
}

func (r *Dingtalker) SendChatTextMessage(msg string) {
	r.Debug("corpSecret=", r.CorpSecret, "\ncorpId=", r.CorpId, "\nchatId=", r.ChatId)
	c := godingtalk.NewDingTalkClient(r.CorpId, r.CorpSecret)
	c.RefreshAccessToken()
	r.Debug("AccessToken = ", c.AccessToken)
	err := c.SendTextMessage("YY", r.ChatId, msg)
	if (err != nil) && (r.Loger != nil) {
		r.Error(err)
	}
}

func (r *Dingtalker) SendRobotTextMessage(msg string) {
	r.Debug("corpSecret=", r.CorpSecret, "\ncorpId=", r.CorpId, "\nchatId=", r.ChatId)
	c := godingtalk.NewDingTalkClient(r.CorpId, r.CorpSecret)
	if c != nil {
		c.RefreshAccessToken()
		r.Debug("AccessToken = ", c.AccessToken)
		err := c.SendRobotTextMessage(r.AcToken, msg)
		if err != nil {
			r.Error(err)
		}
	}
}

func (r *Dingtalker) Error(arg ...interface{}) {
	if r.Loger != nil {
		r.Loger.Error(arg...)
	}
}

func (r *Dingtalker) Debug(arg ...interface{}) {
	if r.Loger != nil {
		r.Loger.Debug(arg...)
	}
}
