// 9. Найдите максимальное из двух входных чисел.

package main;
import "fmt"

func main() {
	var a, b int;
	fmt.Scanln(&a, &b)

	maxNum := max(a, b)
	fmt.Println("Max number is", maxNum)
}