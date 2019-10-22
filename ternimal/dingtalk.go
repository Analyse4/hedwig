package ternimal

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/Analyse4/hedwig/resource"
	"log"
	"net/http"
	"time"
)

type DingTalk struct {
	webHook string
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

func NewDingTalk(wh string) *DingTalk {
	dg := new(DingTalk)
	dg.Text = new(Text)
	dg.At = new(At)
	dg.webHook = wh
	return dg
}

func (dt *DingTalk) Construct(resource *resource.Github) {
	dt.MsgType = "text"
	t := fmt.Sprintf("%v-%v is %v\nlink: %v\n", resource.Repository.Name, resource.Release.TagName, resource.Action, resource.Release.HTMLURL)
	dt.Text.Content = t
}

func (dt *DingTalk) Send() error {
	data, err := json.Marshal(dt)
	if err != nil {
		log.Println(err)
		return err
	}
	buf := bytes.NewBuffer(data)
	//_, err = http.Post(dt.webHook, "application/json", buf)
	req, err := http.NewRequest("POST", dt.webHook, buf)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	tp := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	hClient := new(http.Client)
	hClient.Transport = tp
	_, err = hClient.Do(req)
	if err != nil {
		return err
	}
	log.Printf("%v send to dingtalk robot, text: %v", time.Now(), dt.Text.Content)
	return nil
}
