package main

import "fmt"

func main() {
	multiplicationTable(5)
}

func multiplicationTable(a int) {
	fmt.Print("   ")
	// Вывод заголовков столбцов таблицы
	for i := 1; i <= a; i++ {
		fmt.Printf("%3d ", i)
	}
	// Переход на новую строку
	fmt.Println()

	for i := 1; i <= a; i++ {
		// Вывод номера строки
		fmt.Printf("%2d ", i)
		for j := 1; j <= a; j++ {
			// Вывод значений в строке таблицы умножения
			fmt.Printf("%3d ", i*j)
		}
		// Переход на новую строку после вывода строки таблицы
		fmt.Println()
	}
}
