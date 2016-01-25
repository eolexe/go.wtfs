package main

import (
	"github.com/eolexe/go.lviv.gotchas/typeAFix"
	"github.com/eolexe/go.lviv.gotchas/typeBFix"
)

func main() {
	a := &typeAFix.A{}
	b := typeBFix.NewB(a)

	result := b.GetOrder()
	println(result)
}
