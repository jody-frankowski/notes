// Exercise 1.2: Modify the echo program to print the index and value of each of its arguments, one
// per line.
package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func captureStdout(f func()) string {
	old := os.Stdout // keep backup of the real stdout
	readPipe, writePipe, _ := os.Pipe()
	os.Stdout = writePipe

	channel := make(chan string)

	// Copy the output in a separate goroutine so printing won't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, readPipe)
		channel <- buf.String()
	}()

	f()

	// Back to normal state
	writePipe.Close()
	os.Stdout = old // restoring the real stdout
	out := <-channel

	return out
}

func TestMain(t *testing.T) {
	os.Args = []string{"AAA", "BBB", "CCC"}
	want := "0: AAA\n1: BBB\n2: CCC\n"
	got := captureStdout(main)

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
