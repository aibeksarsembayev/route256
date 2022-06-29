package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var logins int
	fmt.Fscan(reader, &logins) // 1. scan login numbers
	for i := 0; i < logins; i++ {
		var attempts int
		fmt.Fscan(reader, &attempts) // 2. scan for login attemps nubmers

		m := make(map[string]bool)
		loginText := make([]string, attempts)

		for j := 0; j < attempts; j++ {
			fmt.Fscan(reader, &loginText[j])
			loginText[j] = strings.ToLower(loginText[j])
			if !isValid(loginText[j]) {
				fmt.Fprintln(writer, "NO")
			} else {
				if _, ok := m[loginText[j]]; ok {
					fmt.Fprintln(writer, "NO"
				} else {
					fmt.Fprintln(writer, "YES")
					m[loginText[j]] = true
				}
			}
		}

		// fmt.Println(loginText)
		fmt.Fprintln(writer,"")
	}

}

func isValid(s string) bool {
	if len(s) < 2 || 24 < len(s) { //login length 2-24
		return false
	}
	if s[0] == '-' { // should start with '-'
		return false
	}
	for i := 0; i < len(s); i++ {
		if (s[i] < 'a' || s[i] > 'z') && (s[i] < '0' || s[i] > '9') && s[i] != '-' && s[i] != '_' {
			return false
		}
	}
	return true
}
