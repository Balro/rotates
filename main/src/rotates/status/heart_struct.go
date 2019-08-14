package status

import (
	"io"
	"log"
	"net/http"
	"rotates/conf"
	"strconv"
)

func ok(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "ok")
	if err != nil {
		log.Println(err)
	}
}

func StartHeartbeat(c conf.Service) {
	http.HandleFunc("/status", ok)
	err := http.ListenAndServe(":"+strconv.FormatInt(int64(c.Port), 10), nil)
	if err != nil {
		log.Fatal(err)
	}
}
