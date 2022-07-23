package main

import (
	"bufio"
	"fmt"
	"os"
)

func ScanLine() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
	}
	return in.Text()
}

func main() {
	fmt.Println("Введи текст: ")
	t := ScanLine()
	fmt.Println("Введи сдвиг: ")
	var y int
	x := []rune(t)
	fmt.Scan(&y)
	for i := 0; i < len(x); i++ {
		if (x[i] > 96) && (x[i] < 123) {
			x[i] = x[i] + rune(y)
			if x[i] < 97 {
				x[i] = x[i] + rune(26)
			}
			if x[i] > 122 {
				x[i] = x[i] - rune(26)
			}
		}
		if (x[i] > 64) && (x[i] < 91) {
			x[i] = x[i] + rune(y)
			if x[i] < 'A' {
				x[i] = x[i] + rune(26)
			}
			if x[i] > 90 {
				x[i] = x[i] - rune(26)
			}
		}
	}
	fmt.Print(string(x))
}
