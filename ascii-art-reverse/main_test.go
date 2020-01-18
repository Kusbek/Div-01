package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestAsciiArt(t *testing.T) {
	num, answer := "0", "Hello World"
	check(t, "example0"+num, answer)

	num, answer = "1", "123"
	check(t, "example0"+num, answer)

	num, answer = "2", "#=\\["
	check(t, "example0"+num, answer)

	num, answer = "3", "(somthimg&234)"
	check(t, "example0"+num, answer)

	num, answer = "4", "abcdefghijklmnopqrstuvwxyz"
	check(t, "example0"+num, answer)

	num, answer = "5", "\\!\" #$%&'()*+,-./"
	check(t, "example0"+num, answer)

	num, answer = "6", ":;<=>?@"
	check(t, "example0"+num, answer)
	num, answer = "7", "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	check(t, "example0"+num, answer)

}

func check(t *testing.T, input, answer string) {
	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	os.Args = []string{"cmd", "--reverse=" + input + ".txt"}
	// fmt.Println("Testing for argument", os.Args)

	main()

	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	// back to normal state
	w.Close()
	os.Stdout = old // restoring the real stdout
	out := <-outC

	fmt.Println("Testing for argument", input, "and answer", answer)
	if answer+"\n" != string(out) {
		t.Errorf("Ascii art failed, expected\n %v, got\n %v", answer, string(out))
	} else {
		println("PASS")
	}
}
