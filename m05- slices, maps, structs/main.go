package main

import "fmt"

func main() {
	names := []string{"Mateus", "Ramos", "Campeão", "Daora"} // Slice, tem tamanho dinâmico
	fmt.Println(names, len(names), cap(names))
	names = append(names, "Receba")
	fmt.Println(names, len(names), cap(names)) // len -> Tamanho da lista cap -> Capacidade da lista
	names = append(names, "Vitória")
	fmt.Println(names, len(names), cap(names))
	/*
		Ver diferença entre array e slice, link salvo
	*/
	names2 := make([]string, 10, 20) // Vai criar com 10 preenchido vazio e capacidade 20
	// Mas dps se for adicionando com append e passar de 20 ele continua expandidando.
	fmt.Println(names2)
	// MAP -> chave e valor sendo q ambos podem ser de qualquer tipo de dado e diferentes
	idades := make(map[string]uint8)
	idades["Mateus"] = 22
	idades["Ana"] = 21
	fmt.Println(idades)
	// NÃO TEM ORDENAÇÃO, a sequência que vc salva não é confiável dps pra recuperar os valores.
	// Sempre tem q usar a chave certa msm
	// SE buscarmos um indíce que ele não encontra, ele retorna o valor zero daquele tipo de dados
	fmt.Println(idades["Nome Desconhecido"])
	val, ok := idades["Mateus"]
	fmt.Println(val, ok)
	val, ok = idades["Nome desconhecido"]
	fmt.Println(val, ok) // Agora validamos pois o ok veio false, de que o indíce não existe.
	// STRUCTS

	type Pessoa struct {
		Nome      string
		Sobrenome string
		Idade     uint8
		Status    bool
	}
	p := Pessoa{
		Nome:      "Mateus",
		Sobrenome: "Ramos",
		Idade:     22,
		Status:    true,
	}
	fmt.Println(p)
	fmt.Println(p.Nome, p.Sobrenome)

	type Categoria struct {
		Nome string
		pai  *Categoria // Não posso passar a struct dentro da mesma, mas posso passar como ponteiro
	}

	// ATENÇÃO STRUCT Q TEM UMA Propreade começando com letra minúscula é privada
	// Ele consegue ser setado e usado aqui pois está no mesmo package, mas fora não seria possível
}
