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
	check(t, "hello", "hello.txt")
	check(t, "HELLO", "HELLO.txt")
	check(t, "HeLlo HuMaN", "HeLloHuMaN.txt")
	check(t, "1Hello 2There", "1Hello2There.txt")
	check(t, "Hello\nThere", "HelloThere.txt")
	check(t, "{Hello & There #}", "1234.txt")
	check(t, "hello There 1 to 2!", "helloThere1to2.txt")
	check(t, "MaD3IrA&LiSboN", "MaD3IrA&LiSboN.txt")
	check(t, "1a\"#FdwHywR&/()=", "ShittyInput.txt")
	check(t, "{|}~", "SomeRandomSigns.txt")
	check(t, "[\\]^_ 'a", "O.txt")
	check(t, "RGB", "RGB.txt")
	check(t, ":;<=>?@", "Fuck.txt")
	check(t, "\\!\" #$%&'()*+,-./", "lalala.txt")
	check(t, "ABCDEFGHIJKLMNOPQRSTUVWXYZ", "ABCDEFGHIJKLMNOQPRSTUVWXYZ.txt")
	check(t, "abcdefghijklmnopqrstuvwxyz", "abcdefghijklmnopqrstuvwxyz.txt")
}

func check(t *testing.T, input string, fileName string) {
	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	os.Args = []string{"cmd", input}

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
