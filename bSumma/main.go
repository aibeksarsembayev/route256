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

	var counterMain int
	fmt.Fscan(reader, &counterMain)

	for i := 0; i < counterMain; i++ {
		var items int
		fmt.Fscan(reader, &items)

		m := make(map[int]int) // create map to track the same products
		var result int

		for j := 0; j < items; j++ { // loop to receive each product

			var a int
			fmt.Fscan(reader, &a)

			m[a]++ // iterate counter by key for the same products

			if m[a] == 3 { // if it 3th product ..
				delete(m, a) // delete key to eliminate adding to sum int
			} else {
				result += a // otherwise add it to sum
			}

		}
		fmt.Fprintln(writer, result)
	}
}
