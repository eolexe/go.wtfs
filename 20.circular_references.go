package main

import (
	"github.com/eolexe/go.lviv.gotchas/typeA"
	"github.com/eolexe/go.lviv.gotchas/typeB"
)

func main() {
	_ := &typeB.B{}
	_ := &typeA.A{}

}
