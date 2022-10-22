package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type RandomDataGenerator struct{}

func (r RandomDataGenerator) Init() {
	rand.Seed(time.Now().UnixNano())
}

func (r RandomDataGenerator) Generate() string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	id := ""
	for i := 0; i < 100; i++ {
		id += fmt.Sprint(string(chars[rand.Intn(len(chars))]))
	}
	return id
}

func (r RandomDataGenerator) RandomRequestId(length int) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	id := ""
	for i := 0; i < length; i++ {
		id += fmt.Sprint(string(chars[rand.Intn(len(chars))]))
	}
	return strings.ToLower(id)
}
