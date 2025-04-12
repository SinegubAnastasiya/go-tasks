// 8. Напишите программу, которая выводит все нечетные числа от 1 до N.
// Ввод: 10 → Вывод: 1 3 5 7 9
// Ввод: 6 → Вывод: 1 3 5

package main;
import "fmt";

func main() {
	var num int;
	fmt.Scan(&num)

	for i := 1; i <= num; i++ {
		if i % 2 == 1 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}