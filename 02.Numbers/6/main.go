// 6. Напишите программу, которая принимает 3 числа и выводит максимальное из них.
// Ввод: 3, 5, 2 → Вывод: 5
// Ввод: 9, 7, 8 → Вывод: 9

package main;
import "fmt";

func main() {
	var a, b, c int;
	fmt.Scan(&a, &b, &c)

	fmt.Println("Max number is", max(a, b, c))
}