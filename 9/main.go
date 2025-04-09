// 9. Выполните операции с входными числами: Сложение Вычитание Умножение Деление
// Остаток от деления

package main;
import "fmt"

func main() {
	var a, b int;
	fmt.Scanln(&a, &b)

	fmt.Println("Результат сложения: ", a + b)
	fmt.Println("Результат вычитания: ", a - b)
	fmt.Println("Результат умножения: ", a * b)
	fmt.Println("Результат деления: ", a / b)
	fmt.Println("Остаток от деления: ", a % b)
}