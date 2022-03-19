package main

import "fmt"

// métodos são funções basicamente atreladas  à um struct
func (c Categoria) HasParent() bool {
	return c.Pai != nil
}

type Categoria struct {
	Nome string
	Pai  *Categoria
}

type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     uint8
	Status    bool
}

func (p Pessoa) String() string {
	return fmt.Sprintf("Olá, meu nome é %s e eu tenho %d anos", p.Nome, p.Idade)
}

func main() {
	p := Pessoa{"Mateus", "Ramos", 22, true}
	fmt.Println(p)
	cat := Categoria{Nome: "categoria 1", Pai: &Categoria{Nome: "pai"}}
	if !cat.HasParent() {
		fmt.Println("não tem pai")
	}
}

/*
Se vc não passa o ponteiro vc só consegue recuperar os valores do método, se passar o ponteiro
mexe com a referência.
*/
