// 2. Напишите программу, которая выводит сумму чисел от 1 до N, где N вводится пользователем.

package main;
import "fmt";

func main() {
	var n int;
	fmt.Scan(&n);
	sum := 0;

	for i := 1; i <= n; i++ {
		sum += i;
	}
	fmt.Println(sum)
}