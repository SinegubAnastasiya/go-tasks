// 3. Напишите программу, которая по номеру дня недели (от 1 до 7) выводит название дня
// (например, для 1 — "Понедельник").
// Ввод: 2 → Вывод: Вторник
// Ввод: 5 → Вывод: Пятница

package main;
import "fmt";

func main() {
	var n int;
	fmt.Scan(&n)

	switch n {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	}
}