// 6. Объявите переменные разных типов (int, float64, string, bool) и присвойте значения.
// Выведите их.

package main;
import "fmt";

func main() {
	const num int = 5;
	const b float64 = 2.3;
	const str string = "hello";
	const c bool = true;

	fmt.Println(num, b, str, c)
}