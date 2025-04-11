// 1. Напишите программу, которая принимает два числа и выводит наибольшее из них.

package main;
import "fmt";

func main() {
	var a, b int;
	fmt.Scan(&a, &b);

	fmt.Println("Max number is", max(a, b))
}