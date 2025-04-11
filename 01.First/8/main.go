// 8. Введите число с плавающей точкой и преобразуйте его в целое число (округлить).
// Ввод: 3.14 → 3
// Ввод: 7.99 → 8

package main;
import (
	"fmt"
	"math"
)

func main() {
	var number float64;
	fmt.Scan(&number)

	b := math.Round(number)
	fmt.Println((b))
}