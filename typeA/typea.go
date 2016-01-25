package typeA

import "github.com/eolexe/go.lviv.gotchas/typeB"

type A struct {
}

func (a *A) GetCustomer() string {
	b := &typeB.B{}
	return "Customer Max and " + b.GetOrderDetails()
}
