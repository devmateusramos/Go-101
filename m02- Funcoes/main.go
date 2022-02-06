package main

import (
	"fmt"
	"strconv"
)

func hello(nome string) {
	fmt.Println("Hello", nome, "!")
}

func sum(a, b int /*Em vez de a int, b int por ser do msm tipo*/) int {
	return a + b
} /*
Essas funções só podem ser usados nesse mesmo pacote
para usar em outros deve exporta mudando a letra inicial do nome da função
para maiúsculo
*/

func convertAndSum(a int, b string) (total int, err error) {
	i, err := strconv.Atoi(b)
	if err != nil {
		return
	}
	total = a + i
	return
}

func main() {
	hello("Mateus Ramos")

	fmt.Println("Soma:", sum(12, 20))

	total, err := convertAndSum(12, "30")
	fmt.Println("Soma convertida:", total, err)
}
