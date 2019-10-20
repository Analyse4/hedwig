package handler

import (
	"github.com/Analyse4/hedwig/resource"
	"github.com/Analyse4/hedwig/ternimal"
	"io/ioutil"
	"log"
	"net/http"
)

var DINGTALKWEBHOOK = "https://oapi.dingtalk.com/robot/send?access_token=29af91c2d403549dbbda42230c0e915ffb7b24c2c23cba51d9fceb9ced440478"

func ForwardMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
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
