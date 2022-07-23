package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println("Йо, это своеобразный таймер, после ввода времени таймера \nсразу начнётся отсчёт так, что будь готов.\nВведи время на подобии примера: 10h10m10s")
	var x, f string
	fmt.Scan(&x)
	timeNow := time.Now()
	t, err := time.ParseDuration(x)
	if err != nil {
		fmt.Println("ERROR")
		return
	}
	//ttt := time.Now().Add(10 * time.Second)
	for {
		remaining := time.Until(timeNow.Add(t)).Round(time.Second)
		if remaining < 0 {
			break
		}
		fmt.Print("\033[2J\033[H")
		h := int(remaining.Hours())
		remaining -= time.Duration(h) * time.Hour
		m := int(remaining.Minutes())
		remaining -= time.Duration(m) * time.Minute
		s := int(remaining.Seconds())
		f = fmt.Sprintf("%02d:%02d:%02d", h, m, s)
		Numbers(f)
		time.Sleep(time.Second)
	}
	//runfile := exec.Command("date")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Print("\u001B[30;107m")
		} else {
			fmt.Print("\u001B[39;49m")
		}
		fmt.Print("\033[2J\033[H")
		Numbers(f)
		time.Sleep(500 * time.Millisecond)
	}
}

func Numbers(f string) {
	numbers := [][]string{
		strings.Split(
			` $$$$$$\  
$$$ __$$\ 
$$$$\ $$ |
$$\$$\$$ |
$$ \$$$$ |
$$ |\$$$ |
\$$$$$$  /
 \______/ `, "\n"),
		strings.Split(`  $$\   
$$$$ |  
\_$$ |  
  $$ |  
  $$ |  
  $$ |  
$$$$$$\ 
\______|`, "\n"),
		strings.Split(` $$$$$$\  
$$  __$$\ 
\__/  $$ |
 $$$$$$  |
$$  ____/ 
$$ |      
$$$$$$$$\ 
\________|`, "\n"),
		strings.Split(` $$$$$$\  
$$ ___$$\ 
\_/   $$ |
  $$$$$ / 
  \___$$\ 
$$\   $$ |
\$$$$$$  |
 \______/ `, "\n"),
		strings.Split(`$$\   $$\ 
$$ |  $$ |
$$ |  $$ |
$$$$$$$$ |
\_____$$ |
      $$ |
      $$ |
      \__|`, "\n"),
		strings.Split(`$$$$$$$\  
$$  ____| 
$$ |      
$$$$$$$\  
\_____$$\ 
$$\   $$ |
\$$$$$$  |
 \______/ `, "\n"),
		strings.Split(` $$$$$$\  
$$  __$$\ 
$$ /  \__|
$$$$$$$\  
$$  __$$\ 
$$ /  $$ |
 $$$$$$  |
 \______/ `, "\n"),
		strings.Split(`$$$$$$$$\ 
\____$$  |
    $$  / 
   $$  /  
  $$  /   
 $$  /    
$$  /     
\__/      `, "\n"),
		strings.Split(` $$$$$$\  
$$  __$$\ 
$$ /  $$ |
 $$$$$$  |
$$  __$$< 
$$ /  $$ |
\$$$$$$  |
 \______/ `, "\n"),
		strings.Split(` $$$$$$\  
$$  __$$\ 
$$ /  $$ |
\$$$$$$$ |
 \____$$ |
$$\   $$ |
\$$$$$$  |
 \______/ `, "\n"),
		strings.Split(`    
    
$$\ 
\__|
    
$$\ 
\__|
    `, "\n"),
	}
	for i := 0; i < 8; i++ {
		for _, va := range f {
			switch va {
			case '0':
				fmt.Print(numbers[0][i])
			case '1':
				fmt.Print(numbers[1][i])
			case '2':
				fmt.Print(numbers[2][i])
			case '3':
				fmt.Print(numbers[3][i])
			case '4':
				fmt.Print(numbers[4][i])
			case '5':
				fmt.Print(numbers[5][i])
			case '6':
				fmt.Print(numbers[6][i])
			case '7':
				fmt.Print(numbers[7][i])
			case '8':
				fmt.Print(numbers[8][i])
			case '9':
				fmt.Print(numbers[9][i])
			case ':':
				fmt.Print(numbers[10][i])
			}
		}
		fmt.Println("")
	}
}
