package main

import "fmt"

var name string // variavel declarada a nível de package
/*
	Estão disponíveis a qualquer nível de uso dentro de uma função ou outro
escopo dentro do package
*/
var (
	n1 int
	n2 int32
) // Forma de atribuir várias variáveis
// Mas somente esse package main consegue utilizar mesmo em vários arquivos
// Isso pq elas foram declaradas com letra minúscula, fosse maiúscula poderia
func main() {
	var (
		b1 float64
		b2 float32 // essas não são a nível de package, mas da função
	)
	name = "Mateus Ramos"
	b1 = 22
	b2 = 1
	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println("Hello " + name + "!")
	n1 = 23
	n2 = 2
	fmt.Println(n1)
	fmt.Println(n2)
	/*
		Tipos de dados
		Variável => Pequenos espaços na memória que são reservadas em
		tempo de execução e recuperar dps para fazer algo
	*/
	var total int
	fmt.Println(total)
	// variáveis declaradas tem um valor inicial já passada pra elas como
	// 0 pra int, string vazia pra string etc
	total++
	fmt.Println(total)

	var e, f, g int32 // pode declarar de uma vez assim por serem
	// do mesmo tipo
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	// atribuição de valor em Go possível
	var x, y = 10, 20
	fmt.Printf("x: %v\n", x)
	fmt.Printf("y: %v\n", y)
	x, y = y, x
	fmt.Printf("x: %v\n", x)
	fmt.Printf("y: %v\n", y)
}

// Complementar https://aprendagolang.com.br/2021/11/18/entendendo-os-tipos-de-dados-e-suas-capacidades/?utm_source=youtube&utm_medium=description&utm_campaign=go101
// e https://aprendagolang.com.br/2021/11/22/convertendo-tipos-de-dados/?utm_source=youtube&utm_medium=description&utm_campaign=go101
