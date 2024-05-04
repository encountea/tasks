package main

import "fmt"

func main() {
	fmt.Println(primeNumbers(1, 5))
}

func primeNumbers(a, b int) []int {
	// Создание мапы для хранения делителей и их количества
	counter2 := make(map[int]int)
	// Создание слайса для хранения простых чисел
	var array2 []int
	// Проходим циклом в диапазоне от a до b
	for i := a; i <= b; i++ {
		// Проверка наличия делителей для числа
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				// Увелечение счетчика для делителя
				counter2[i]++
			}
		}
	}
	// Проверка, является ли число простым, т. е. имеет всего 2 делителя
	for m, k := range counter2 {
		if k == 2 {
			// Если число простое, то добавляем его в результирующий слайс
			array2 = append(array2, m)
		}
	}
	return array2
}
