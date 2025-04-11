// 11. Вычислите среднее арифметическое двух чисел.
// Вход: 10 и 20 → Среднее = 15

package main;
import "fmt"

func main() {
	var a, b, avg float32;
	fmt.Scanln(&a, &b)

	avg = (a + b) / 2;
	fmt.Println(avg)
}