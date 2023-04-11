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

func WriteLine(content ...any) {
	fmt.Println(content...)
}

func Write(content ...any) {
	fmt.Print(content...)
}

func ErrWriteLine(content ...any) {
	fmt.Fprintln(os.Stderr, content...)
}

func ErrWrite(content ...any) {
	fmt.Fprint(os.Stderr, content...)
}

func WriteLineF(format string, params ...any) {
	fmt.Printf(format, params...)
	fmt.Println("")
}

func WriteF(format string, params ...any) {
	fmt.Printf(format, params...)
}

func ErrWriteLineF(format string, params ...any) {
	fmt.Fprintf(os.Stderr, format, params...)
	fmt.Println("")
}

func ErrWriteF(format string, params ...any) {
	fmt.Fprintf(os.Stderr, format, params...)
}
