// 12. Сравнение двух входных чисел
// Проверьте: Если a > b — "a больше" Если a == b — "a равно b" Если a < b — "b больше«

package main;
import "fmt"

func main() {
	var a, b int;
	fmt.Scanln(&a, &b)

	if a > b {
		fmt.Println("a больше")
	} else if a == b {
		fmt.Println("a равно b")
	} else {
		fmt.Println("b больше")
	}
}