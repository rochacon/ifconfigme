package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", ifconfig)
	log.Println("Listening on " + os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

func ifconfig(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.Header['X-Forwarded-For'][0]))
}
