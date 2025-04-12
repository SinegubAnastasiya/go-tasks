// 7. Напишите программу, которая вычисляет факториал числа.
// Ввод: 5 → Вывод: 120
// Ввод: 7 → Вывод: 5040

package main;
import "fmt";

func main() {
	var num int;
	fmt.Scan(&num)
	fact := 1;

	for i := 1; i <= num; i++ {
		fact *= i
	}
	fmt.Println(fact)
}