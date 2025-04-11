// 2. Оцените успеваемость входных данных:
// 90-100 — "Отлично".
// 75-89 — "Хорошо".
// 50-74 — "Удовлетворительно".
// < 50 — "Неудовлетворительно".

package main;
import "fmt";

func main() {
	var data int;
	fmt.Scan(&data)

	if data >= 0 && data < 50 {
		fmt.Println("Неудовлетворительно")
	} else if data >=50 && data <=74 {
		fmt.Println("Удовлетворительно")
	} else if data >=75 && data <=89 {
		fmt.Println("Хорошо")
	} else if data >=90 && data <=100 {
		fmt.Println("Отлично")
	} else {
		fmt.Println("Некорректный ввод")
	}
}