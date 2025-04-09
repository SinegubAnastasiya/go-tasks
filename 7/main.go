// 7. Преобразуйте целое число в строку и строку в число.
// Вход: 123 → "123" (преобразовать в строку)
// Вход: "456" → 456 (преобразовать в число)

package main;
import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 10;
	var str string = "5";

	fmt.Println(strconv.Itoa(num))

	strToNum, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(strToNum)
	}
}