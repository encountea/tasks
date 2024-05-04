package main

import "fmt"


func main() {
	nums := []int{42, 12, 18}
	fmt.Println(divider(nums))
}

func divider(nums []int) []int {
	// Создание мапы для хранения делителей и их количества
	counter := make(map[int]int)
	// Создание слайса для хранения общих делителей
	var array []int

	for _, v := range nums {
		// Проверка наличия делителей для числа
		for i := 2; i <= v; i++ {
			if v % i == 0 {
				// Увелечение счетчика для делителя
				counter[i]++
			}
		}
	}

	// Проверка, является ли делитель общим для всех чисел в слайсе
	for j, k := range counter {
		if k == len(nums) {
			// Если делитель общий для всех чисел, добавляем его в результирующий слайс
			array = append(array, j)
		}
	}
	return array
}