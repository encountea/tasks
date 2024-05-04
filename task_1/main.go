package main

import "fmt"

func main() {
    fmt.Println(computers(1121))
}

func computers(a int) string {
	if  a % 10 == 1 && a % 100 != 11 && a % 100 != 12 && a % 100 != 13 && a % 100 != 14 {
        return fmt.Sprintf("%d компьютер", a)
    } else if  a % 10 == 5 || a % 10 == 6 || a % 10 == 7 || a % 10 == 8 || a % 10 == 9 || a % 10 == 0 {
        return fmt.Sprintf("%d компьютеров", a)
    } else if  (a > 10 && a < 15) || a > 100 && (a % 100 == 11 || a % 100 == 12 || a % 100 == 13 || a % 100 == 14) {
        return fmt.Sprintf("%d компьютеров", a)
    } else {
        return fmt.Sprintf("%d компьютера", a)
    }
}