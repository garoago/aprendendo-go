// Package main provides ...
package main

import "fmt"

func main() {

	var foo int
	for foo = 0; foo < 10; foo++ {
	}
	fmt.Println("Usando variavel global: ", foo)

	// com contexto
	var bar int
	for bar := 0; bar < 10; bar++ {
	}
	fmt.Println("Entendeu agora o ':='? E sim, contexto: ", bar)

}
