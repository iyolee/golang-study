package main

import (
	"fmt"
	"net/http"
	"web"
)

func main() {
	http.HandleFunc("/test", web.Echo)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
