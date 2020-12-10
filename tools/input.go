package tools

import (
	"fmt"
	"strings"
)

// Input validates the Input and returns the coordinates
func Input() (int, int) {
	var x, y int
	for {
		var input string

		fmt.Print("Enter Position: ")
		fmt.Scanln(&input)

		if strings.Contains("ABC", strings.ToUpper(string(input[0]))) && strings.Contains("123", string(input[1])) {
			x = int(input[0]) - 97
			y = int(input[1]) - 49
			break
		} else {
			fmt.Println("Invalid Input")
		}
	}
	return x, y
}
