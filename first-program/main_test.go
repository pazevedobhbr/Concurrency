package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

/*func Test_updateMessage(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	var wg sync.WaitGroup
	wg.Add(1)
	go updateMessage("universe", &wg)
	wg.Wait()
	printMessage()
	_ = w.Close()
	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "Hello, universe") {
		t.Errorf("UpdateSomething() = %v, want %v", output, "Hello, universe")
	}
}*/

func Test_updateMessage(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	go updateMessage("epsilon", &wg)

	wg.Wait()

	if msg != "Hello, epsilon" {
		t.Error("Expected to find epsilon, but it is not there")
	}
}

func Test_printMessage(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	msg = "epsilon"
	printMessage()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "epsilon") {
		t.Error("Expected to find epsilon, but it is not there")
	}
}

func Test_main(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "Hello, universe!") {
		t.Error("Expected to find Hello, universe!, but it is not there")
	}

	if !strings.Contains(output, "Hello, cosmos!") {
		t.Error("Expected to find Hello, cosmos!, but it is not there")
	}

	if !strings.Contains(output, "Hello, world!") {
		t.Error("Expected to find Hello, world!, but it is not there")
	}
}
