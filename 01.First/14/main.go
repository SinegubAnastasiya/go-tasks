// 10. Найдите минимальное из трех входных чисел.

package main;
import "fmt"

func main() {
	var a, b, c int;
	fmt.Scanln(&a, &b, &c)

	minNum := min(a, b, c)
	fmt.Println("Min number is", minNum)
}