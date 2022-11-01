package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	depSeq := os.Getenv("DEPLOYMENT_SEQUENCE")
	rdg := RandomDataGenerator{}
	rdg.Init()
	reqId := rdg.RandomRequestId(10)
	data := rdg.Generate()
	response := ""

	response += fmt.Sprintln("---------------------------")
	response += fmt.Sprintln("Deployment Sequence: ", depSeq)
	response += fmt.Sprintln("Request ID: ", reqId)
	response += fmt.Sprintln("Data: ", data)
	response += fmt.Sprintln("---------------------------")

	w.Header().Add("Content-Type", "text/plain")
	w.Write([]byte(response))
}

func main() {
	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "3200"
	}
	http.HandleFunc("/", handler)
	fmt.Println("Listening on port ", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		fmt.Println(fmt.Errorf("server error: %v", err))
	}
}
