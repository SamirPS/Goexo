// You can edit this code!
// Click here and start typing.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func Palindrome(s string) string {
	for i := range s {
		if s[i] != s[len(s)-i-1] {
			return "non"
		}
	}
	return "oui"
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Printf("Le mot %s est-il un Palindrome ? %s \n", scanner.Text(), Palindrome(scanner.Text()))
	}

}
