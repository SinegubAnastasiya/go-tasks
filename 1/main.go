// 1. Проверьте возраст у входных данных :
// Если < 18 — "Вы несовершеннолетний".
// Если 18-60 — "Вы взрослый".
// Если > 60 — "Вы пожилой человек"

package main;
import "fmt";

func main() {
	var age int;
	fmt.Scan(&age)
	if age < 0 {
		fmt.Println("Некорректно введен возраст")
	} else if age < 18 {
		fmt.Println("Вы несовершеннолетний")
	} else if age >= 18 && age <= 60 {
		fmt.Println("Вы взрослый")
	} else {
		fmt.Println("Вы пожилой человек")
	}
}