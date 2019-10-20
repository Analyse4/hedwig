package handler

import (
	"github.com/Analyse4/hedwig/resource"
	"github.com/Analyse4/hedwig/ternimal"
	"io/ioutil"
	"log"
	"net/http"
)

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
		log.Println(digest)
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
	dgt := ternimal.NewDingTalk()
	err = dgt.Construct(ghr)
	//send
}
