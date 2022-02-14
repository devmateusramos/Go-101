package main

import (
	"fmt"
	"log"
	"os"
)

var (
	cara, coroa int
)

func lancarMoeda(lado string) {
	switch lado {
	case "cara":
		cara++

	case "coroa":
		coroa++

	default:
		fmt.Println("Caiu em pé")
	}
}

func main() {
	a, b := 5, 10

	if a > b {
		fmt.Println("A é maior do que B")
	} else if a < b {
		fmt.Println("A é menor do que B")
	} else {
		fmt.Println(" A e B são iguais")
	}
	// refazendo em switch case
	switch {
	case a > b:
		fmt.Println("A é maior do que B")

	case a < b:
		fmt.Println("A é menor do que B")

	default:
		fmt.Println("A e B são iguais")
	}

	file, err := os.Open("hello.txt") // aqui o err existe
	// para toda aplicação
	if err != nil {
		log.Panic(err)
	}

	data := make([]byte, 100)
	if _, err := file.Read(data); err != nil { // aqui o err só
		// existe dentro do if
		log.Panic(err)
	}

	fmt.Println(string(data))

}
