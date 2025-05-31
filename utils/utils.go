package utils

import (
	"bufio"
	"fmt"
	"os"
)

func GracefulExit() {
	fmt.Println("\nPress any key to exit")
	bufio.NewReader(os.Stdin).ReadLine()
	os.Exit(1)
}
