package belajar_golang_context

import (
	"context"
	"fmt"
	"testing"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
	// ContextA sebagai Parent
	contextA := context.Background()

	// contextB dan C mempunyai parent ContextA
	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	// context D dan E mempunyai parent ContextB
	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")
	contextG := context.WithValue(contextF, "g", "G")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)
	fmt.Println(contextG)

	// kode context get value, ini akan bertanya ke parent bukan ke childnya.
	fmt.Println("Context F:", contextF.Value("f"))
	fmt.Println("Context F:", contextF.Value("c"))
	fmt.Println("Context F:", contextF.Value("b")) // beda parent, maka tidak dapat
	fmt.Println("Context A:", contextA.Value("b"))
}
