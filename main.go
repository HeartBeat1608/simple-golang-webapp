package main

import (
	"fmt"
)

func main() {
	rdg := RandomDataGenerator{}
	rdg.Init()
	reqId := rdg.RandomRequestId(10)
	data := rdg.Generate()

	fmt.Println("---------------------------")
	fmt.Println("Request ID: ", reqId)
	fmt.Println("Data: ", data)
	fmt.Println("---------------------------")
}
