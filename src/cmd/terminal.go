package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func ReadInput(message string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf(message)

	input, _ := reader.ReadString('\n')

	return strings.Replace(input, "\n", "", -1)
}
