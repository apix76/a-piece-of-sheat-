package main

import (
	"fmt"
)

func zam(x, y []int) {
	b, a := 0, len(y)-1
	for {
		if (x[b] != 0) && (y[a] == 0) && ((a+1 == len(y)) || (x[b] < y[a+1])) {
			y[a] = x[b]
			x[b] = 0
			break
		}
		if x[b] == 0 {
			b++
		}
		if y[a] != 0 {
			a--
		}
		if (a < 0) || (b > len(x)-1) {
			break
		}
		if a == len(y)-1 {
			continue
		}
		if x[b] > y[a+1] {
			break
		}
	}
}

func prn(x, y, z []int) {
	for i := 0; i < len(x); i++ {
		fmt.Println(x[i], y[i], z[i])
	}
}

func main() {
	count, g, a := 0, 0, 0
	var x, y, z []int
	var f, h string
	fmt.Println("Введи количество дисков")
	fmt.Scan(&g)
	for i := 0; i < g; i++ {
		x, y, z, f = append(x, i+1), append(y, 0), append(z, 0), f+string(i+1)
	}
	prn(x, y, z)
	for {
		if h == f {
			break
		}
		h = ""
		fmt.Println("Слитно введите номер первичного стобца и вторичного:")
		fmt.Scan(&a)
		switch a {
		case 12:
			zam(x, y)
		case 13:
			zam(x, z)
		case 21:
			zam(y, x)
		case 23:
			zam(y, z)
		case 31:
			zam(z, x)
		case 32:
			zam(z, y)
		default:
			fmt.Print("Ошибка ввода.")
		}
		for _, va := range z {
			h = h + string(va)
		}
		fmt.Print("\033[2J\033[H")
		prn(x, y, z)
		count += 1
		fmt.Println("Количество шагов: ", count)
	}
	fmt.Println("Капитан залупа с корабля сбежал!")
}
