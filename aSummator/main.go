package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	readerInput := bufio.NewReader(os.Stdin)
	writerOutput := bufio.NewWriter(os.Stdout)
	defer writerOutput.Flush()

	var counter int
	fmt.Fscan(readerInput, &counter)

	for counter > 0 {
		var a, b int
		fmt.Fscan(readerInput, &a, &b)
		fmt.Fprintln(writerOutput, a+b)
		counter--
	}
}
