package typeB

import "github.com/eolexe/go.lviv.gotchas/typeA"

type B struct {
}

func (b *B) GetOrderDetails() string {
	return "Order details: 2 jars of peanuts"
}

func (b *B) GetOrder() string {
	a := typeA.A{}
	return a.GetCustomer()
}
