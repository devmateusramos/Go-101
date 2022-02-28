package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
	nomes := []string{"Mateus", "Ramos", "Vencendor"}
	for j := 0; j < len(nomes); j++ {
		fmt.Println(nomes[j])
	}

	for k, nome := range nomes {
		fmt.Println(k, nome)
	}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	var i int

	for i < len(nomes) {
		fmt.Println(nomes[i] + " sim")
		i++
	}
	/*
		Outro uso do for é para um looping infinito como por exemplo abrir um server
		q vc quer fique escutando sem parar
		for {
			fmt.Println(nomes[i])
			i++
		}
	*/
}
