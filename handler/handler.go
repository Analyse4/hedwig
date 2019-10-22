package handler

import (
	"github.com/Analyse4/hedwig/resource"
	"github.com/Analyse4/hedwig/ternimal"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var DINGTALKWEBHOOK = "https://oapi.dingtalk.com/robot/send?access_token=a7cecb1ead549ecc0f94187e69753271f8d389f2b23972fa59b94e5af78f028c"

func ForwardMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	event := r.Header.Get("X-GitHub-Event")
	if event == "ping" {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	digest := r.Header.Get("X-Hub-Signature")
	if digest != "" {
		// verify body integrity
		log.Printf("digest: %v\n", digest)
	}
	log.Printf("%v recieve event: %v", time.Now(), event)
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	ghr := resource.NewGithub()
	err = ghr.Construct(body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	dgt := ternimal.NewDingTalk(DINGTALKWEBHOOK)
	dgt.Construct(ghr)
	err = dgt.Send()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
