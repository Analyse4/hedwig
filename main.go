package hedwig

import (
	"github.com/Analyse4/hedwig/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.ForwardMessage)
	log.Fatal(http.ListenAndServe(":2245", nil))
}
