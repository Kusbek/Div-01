package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestAsciiOutput(t *testing.T) {
	// tester(t, "First\nTest", "shadow", "test00")
	tester(t, "hello", "standard", "hello")
	tester(t, "123 -> #$%", "standard", "test02.txt")
	tester(t, "432 -> #$%&@", "shadow", "test03")
	tester(t, "There", "shadow", "test04.txt")
	tester(t, "123 -> \"#$%@", "thinkertoy", "test05.txt")
	tester(t, "2 you", "thinkertoy", "test06.txt")
	tester(t, "Testing long output!", "standard", "test07.txt")
}

func tester(t *testing.T, input, option, fileName string) {
	os.Args = []string{"cmd", input, option, "--output=" + fileName}
	main()
	output, _ := ioutil.ReadFile(fileName + ".txt")
	answer, _ := ioutil.ReadFile(fileName + "_test.txt")
	fmt.Println("Testing for argument", input)
	if string(answer) != string(output) {
		t.Errorf("Ascii art output failed, expected\n %v, got\n %v", string(answer), string(output))
	} else {
		println("PASS")
	}
}
