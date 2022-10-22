package main

import (
	"testing"
)

func TestRandomDataGenerator(t *testing.T) {
	rdg := RandomDataGenerator{}
	rdg.Init()
	reqId := rdg.RandomRequestId(10)
	data := rdg.Generate()

	if len(reqId) != 10 {
		t.Errorf("Expected length of 10, got %d", len(reqId))
	}

	if len(data) != 100 {
		t.Errorf("Expected length of 100, got %d", len(data))
	}
}

func TestRandomRequestId(t *testing.T) {
	rdg := RandomDataGenerator{}
	rdg.Init()
	reqId := rdg.RandomRequestId(10)

	if len(reqId) != 10 {
		t.Errorf("Expected length of 10, got %d", len(reqId))
	}
}

func TestGenerate(t *testing.T) {
	rdg := RandomDataGenerator{}
	rdg.Init()
	data := rdg.Generate()

	if len(data) != 100 {
		t.Errorf("Expected length of 100, got %d", len(data))
	}
}
