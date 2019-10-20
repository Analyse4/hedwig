package ternimal

import "github.com/Analyse4/hedwig/resource"

type DingTalk struct {
	MsgType string `json:"msgtype"`
	Text    *Text  `json:"text"`
	At      *At    `json:"at"`
}
type Text struct {
	Content string `json:"content"`
}
type At struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll   bool     `json:"isAtAll"`
}

func NewDingTalk() *DingTalk {
	dg := new(DingTalk)
	dg.Text = new(Text)
	dg.At = new(At)
	return dg
}

func (dt *DingTalk) Construct(resource *resource.Github) error {
	dt.MsgType = "text"
	// construct text
	return nil
}
