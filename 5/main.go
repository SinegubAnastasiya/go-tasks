// 5. Проверьте, имеет ли человек право на льготу (меньше 18 или старше 65 лет).
// Вход: 15 → "Льгота есть"
// Вход: 40 → "Льгота нет"
// Вход: 70 → "Льгота есть«

package main;
import "fmt";

func main() {
	var age int;
	fmt.Scan(&age)

	if age < 0 {
		fmt.Println("Некорректно введен возраст")
	} else if age < 18 {
		fmt.Println("Льгота есть")
	} else if age >= 18 && age <= 65 {
		fmt.Println("Льгота нет")
	} else {
		fmt.Println("Льгота есть")
	}
}