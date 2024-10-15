package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	separator := scanner.Text()
	_ = scanner.Scan()
	string1 := scanner.Text()
	_ = scanner.Scan()
	string2 := scanner.Text()
	_ = scanner.Scan()
	string3 := scanner.Text()
	fmt.Print(string1, separator, string2, separator, string3)
}
