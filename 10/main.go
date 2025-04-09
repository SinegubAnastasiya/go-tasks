// 10. Вычислите периметр прямоугольника по длине и ширине.
// Вход: длина = 5, ширина = 3 → Периметр = 16

package main;
import "fmt"

func main() {
	var a, b int;
	fmt.Scanln(&a, &b)

	p := (a + b) * 2;
	fmt.Println(p)
}