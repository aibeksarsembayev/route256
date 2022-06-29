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
		// fmt.Fprintln(writer, tabInt)

		var clicks int
		fmt.Fscan(reader, &clicks) // 4. scan clicks numbers

		clickCol := make([]int, clicks)
		for n := 0; n < clicks; n++ {
			fmt.Fscan(reader, &clickCol[n]) // 5. scan click column to clicks slice
		}
		// fmt.Fprintln(writer, clickCol)

		result := sortByCol(tabInt, clickCol, row, col)
		// fmt.Fprintln(writer, result)

		printResult(writer, result, row, col)
	}
}

func sortByCol(tabInt [][]int, clickCol []int, row, col int) [][]int {
	for i := 0; i < len(clickCol); i++ {
		colNum := clickCol[i] - 1
		for j := 0; j < row-1; j++ {
			for k := 0; k < row-1; k++ {
				if tabInt[k][colNum] > tabInt[k+1][colNum] {
					tabInt[k][colNum], tabInt[k+1][colNum] = tabInt[k+1][colNum], tabInt[k][colNum]
					for m := 0; m < col; m++ {
						if m != colNum {
							tabInt[k][m], tabInt[k+1][m] = tabInt[k+1][m], tabInt[k][m]
						}
					}
				}
			}
		}
	}
	return tabInt
}

func printResult(writer *bufio.Writer, tabInt [][]int, row, col int) {
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Fprint(writer, tabInt[i][j])
			if j == col-1 {
				fmt.Fprintln(writer, "")
			} else {
				fmt.Fprint(writer, " ")
			}
		}
	}
	fmt.Fprintln(writer, "")
}
