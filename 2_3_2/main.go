package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin) // создаем копию структуры bufio.Scanner
	_ = scanner.Scan()                    // на этом месте приложение останавливается и ожидает ввода. Завершением ввода, будет нажатие Enter
	book := scanner.Text()                // cохраняем все что ввели в переменную "name"
	fmt.Println(book, "- лучшая книга!")
}
