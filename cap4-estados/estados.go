package main

import "fmt"

type Estado struct {
	nome      string
	populacao int
	capital   string
}

func main() {
	estados := make(map[string]Estado, 6)

	estados["GO"] = Estado{"Goiás", 6434052, "Goiania"}
	estados["PB"] = Estado{"Paraíba", 3914418, "João Pessoa"}
	estados["PR"] = Estado{"Paraná", 10997462, "Curitiba"}

	fmt.Println(estados)
}
