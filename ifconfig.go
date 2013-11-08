package main

import (
	"log"
	"net/http"
	"os"
	//"strings"
)

func main() {
	http.HandleFunc("/", ifconfig)
	log.Println("Listening on " + os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

func ifconfig(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		if len(v) > 0 {
            w.Write([]byte(k+":"+v[0]+"\n"))
		}
	}
}
