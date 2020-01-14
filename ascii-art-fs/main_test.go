package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestAsciiArt(t *testing.T) {
	check(t, "hello", "standard", "hello.txt")
	check(t, "hello world", "shadow", "HELLO.txt")
	check(t, "nice 2 meet you", "thinkertoy", "nicetomeetyou.txt")
	check(t, "you & me", "standard", "youandme.txt")
	check(t, "123", "shadow", "123.txt")
	check(t, "/(\")", "thinkertoy", "shit.txt")
	check(t, "ABCDEFGHIJKLMNOPQRSTUVWXYZ", "shadow", "ABCDEFGHIJKLMNOPQRSTUVWXYZ.txt")
	check(t, "\"#$%&/()*+,-./", "thinkertoy", "shit2.txt")
	check(t, "It's Working", "thinkertoy", "itsworking.txt")

}

func check(t *testing.T, input string, option string, fileName string) {
	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	os.Args = []string{"cmd", input, option}

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

	output, _ := ioutil.ReadFile(fileName)

	fmt.Println("Testing for argument", input)
	if out != string(output) {
		// for i := range output {
		// 	if out[i] != output[i] {
		// 		if output[i] != 10 {
		// 			output[i] = 42
		// 		}
		// 	}
		// }
		t.Errorf("Ascii art failed, expected\n %v, got\n %v", string(output), out)
	} else {
		println("PASS")
	}
}
