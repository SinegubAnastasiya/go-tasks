// 5. Напишите программу, которая проверяет, делится ли введенное число на 5 и 10.
// Ввод: 25 → Вывод: Делится на 5, но не на 10
// Ввод: 30 → Вывод: Делится на 5 и на 10

package main;
import "fmt";

func main() {
	var n int;
	fmt.Scan(&n)

	if n % 5 == 0 && n % 10 == 0 {  
		fmt.Println("Делится на 5 и на 10")
	} else if n % 5 == 0 && n % 10 != 0 {
		fmt.Println("Делится на 5, но не на 10")
	} else {
		fmt.Println("Не делится ни на 5, ни на 10")
	}
}