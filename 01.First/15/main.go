// 11. Проверьте, равны ли два входных числа

package main;
import "fmt"

func main() {
	var a, b int;
	fmt.Scanln(&a, &b)

	if a == b {
		fmt.Println("a равно b")
	} else {
		fmt.Println("a не равно b")
	}
}