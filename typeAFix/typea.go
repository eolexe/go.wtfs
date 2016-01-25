package typeAFix

import "github.com/eolexe/go.lviv.gotchas/typeBFix"

type A struct {
}

func (a *A) GetCustomer() string {
	b := &typeBFix.B{}
	return "Customer Max and " + b.GetOrderDetails()
}
