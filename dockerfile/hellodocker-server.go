package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	port := 80
	if len(os.Args) > 1 {
		port, _ = strconv.Atoi(os.Args[1])
	}
	addr := ":" + strconv.Itoa(port)

	// Set routing rules
	http.HandleFunc("/", Tmp)

	log.Println("starting on " + strconv.Itoa(port))

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func Tmp(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello docker!")
}
