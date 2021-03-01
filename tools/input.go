package tools

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Input validates the Input and returns the coordinates
func Input() (int, int) {
	var x, y int
	for {
		var input string

		fmt.Print("Enter Position: ")
		scanner := bufio.NewScanner(os.Stdin) // using *bufio.Scanner because fmt.Scan() does not accept empty inputs
		if scanner.Scan() {
			input = scanner.Text()
		}
		if len(input) != 2 || !(strings.Contains("ABC", strings.ToUpper(string(input[0]))) && strings.Contains("123", string(input[1]))) {
			fmt.Println("Invalid Input")
		} else {
			x = int(input[0]) - 97
			y = int(input[1]) - 49
			break
		}
	}
	return x, y
}
