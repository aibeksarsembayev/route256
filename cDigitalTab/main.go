package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tabs int
	fmt.Fscan(reader, &tabs) // 1. scan for overall tabs
	for i := 0; i < tabs; i++ {

		var row, col int
		fmt.Fscan(reader, &row, &col) // 2. scan for tabs dimensions

		tabInt := make([][]int, row)
		for j := range tabInt {
			tabInt[j] = make([]int, col)
		}

		for k := 0; k < row; k++ {
			for l := 0; l < col; l++ {
				fmt.Fscan(reader, &tabInt[k][l]) // 3. scan tabs to 2d slice
			}
		}
		fmt.Fprintln(writer, tabInt)

		var clicks int
		fmt.Fscan(reader, &clicks) // 4. scan clicks numbers

		clickCol := make([]int, clicks)
		for n := 0; n < clicks; n++ {
			fmt.Fscan(reader, &clickCol[n]) // 5. scan click column to clicks slice
		}
		fmt.Fprintln(writer, clickCol)
	}
}
