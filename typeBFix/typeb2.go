package typeBFix

type B struct {
	aType IOrderable
}

func (b *B) GetOrderDetails() string {
	return "Order details: 2 jars of peanuts"
}

func (b *B) GetOrder() string {
	return b.aType.GetCustomer()
}

type IOrderable interface {
	GetCustomer() string
}

func NewB(inject IOrderable) *B {
	return &B{
		aType: inject,
	}
}
