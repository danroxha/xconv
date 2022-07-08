package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadInput(message string, args ...any) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf(message, args...)

	input, _ := reader.ReadString('\n')

	return strings.TrimSpace(strings.Replace(input, "\n", "", -1))
}
