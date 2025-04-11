// 4. Напишите программу, которая принимает число N и выводит все числа от 1 до N, делящиеся
// на 3
// Ввод: 10 → Вывод: 3 6 9
// Ввод: 20 → Вывод: 3 6 9 12 15 18

package main;
import "fmt";

func main() {
	var n int;
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		if i % 3 == 0 {
			fmt.Print(i, " ")
		}
	}	
	fmt.Println()
}