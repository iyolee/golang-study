package web

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Echo server
func Echo(w http.ResponseWriter, r *http.Request) {
	msg, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("echo error"))
		return
	}

	writenLen, err := w.Write(msg)
	if err != nil || writenLen != len(msg) {
		fmt.Println(err, "write len:", writenLen)
	}
}
