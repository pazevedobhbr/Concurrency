package main

import (
	"testing"
	"time"
)

func Test_dine(t *testing.T) {
	eatTime = 0 * time.Second
	sleepTime = 0 * time.Second
	thinkTime = 0 * time.Second

	for i := 0; i < 10; i++ {
		orderFinished = []string{}
		dine()
		if len(orderFinished) != 5 {
			t.Errorf("Incorrect Lengthof Slice. Expected 5 philosophers to finish, but got %d", len(orderFinished))
		}
	}
}

func Test_dineWithVaryingDelays(t *testing.T) {
	var theTests = []struct {
		name  string
		delay time.Duration
	}{
		{"zero Delay", 0 * time.Second},
		{"quarter second delay", 0 * time.Millisecond * 250},
		{"half second delay", 0 * time.Millisecond * 500},
	}
	for _, e := range theTests {
		orderFinished = []string{}
		eatTime = e.delay
		thinkTime = e.delay
		sleepTime = e.delay
		dine()
		if len(orderFinished) != 5 {
			t.Errorf("%s:Incorrect Lengthof Slice. Expected 5 philosophers to finish, but got %d", e.name, len(orderFinished))
		}
	}
}
