// 4. Проверьте, является ли число четным или нечетным.

package main;
import "fmt";

func main() {
	var num int;
	fmt.Scan(&num);

	if num % 2 == 0 {
		fmt.Println("even number")
	} else {
		fmt.Println("odd number")
	}
}