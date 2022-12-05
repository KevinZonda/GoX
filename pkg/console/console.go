package console

import (
	"bufio"
	"fmt"
	"os"
)

var scanner = bufio.NewReader(os.Stdin)

func ReadLine() (string, error) {
	return scanner.ReadString('\n')
}

func WriteLine(format string, params ...any) {
	fmt.Printf(format, params...)
	fmt.Println("")
}

func Write(format string, params ...any) {
	fmt.Printf(format, params...)
}

func ErrWriteLine(format string, params ...any) {
	fmt.Fprintf(os.Stderr, format, params...)
	fmt.Println("")
}

func ErrWrite(format string, params ...any) {
	fmt.Fprintf(os.Stderr, format, params...)
}
