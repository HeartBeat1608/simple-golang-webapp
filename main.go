package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	rdg := RandomDataGenerator{}
	rdg.Init()
	reqId := rdg.RandomRequestId(10)
	data := rdg.Generate()
	response := ""

	response += fmt.Sprintln("---------------------------")
	response += fmt.Sprintln("Request ID: ", reqId)
	response += fmt.Sprintln("Data: ", data)
	response += fmt.Sprintln("---------------------------")

	w.Header().Add("Content-Type", "text/plain")
	w.Write([]byte(response))
}

func main() {
	http.HandleFunc("/", handler)

	done := make(chan bool)

	go func() {
		if err := http.ListenAndServe(":8000", nil); err != nil {
			done <- true
		}
	}()

	<-done
	fmt.Println("Server shutting down")
}
