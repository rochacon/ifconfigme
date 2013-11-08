package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", ifconfig)
	log.Println("Listening on " + os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

func ifconfig(w http.ResponseWriter, r *http.Request) {
	xff := r.Header["X-Forwarded-For"]
	if len(xff) > 0 {
		io.WriteString(w, xff[0])
	} else {
		io.WriteString(w, strings.Split(r.RemoteAddr, ":")[0])
	}
}
