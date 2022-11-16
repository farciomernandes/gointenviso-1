package entity

import "errors"

type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

// Como se fosse um construtor. Retorna 2 valores
func NewOrder(id string, price float64, tax float64) (*Order, error) {
	order := &Order{
		ID:    id,
		Price: price,
		Tax:   tax,
	}

	err := order.isValid()
	if err != nil {
		return nil, err
	}
	return order, nil
}

// (o *Order) = Avisa que este método pode ser acessado a partir de um Order
func (o *Order) isValid() error {
	if o.ID == "" {
		return errors.New("invalid id")
	}
	if o.Price <= 0 {
		return errors.New("invalid price")
	}
	if o.Tax <= 0 {
		return errors.New("invalid Tax")
	}
	return nil
}

/*
Quando uma função tem o (x *Struct) significa que ele é um método da struct criada

No caso do NewOrder funciona assim: order := NewOrder(...)
No caso do isValid funciona assim: order := NewOrder(...);  order.isValid()
*/
func (o *Order) CalculateFinalPrice() error {
	o.FinalPrice = o.Price + o.Tax
	err := o.isValid()

	if err != nil {
		return err
	}
	return nil
}
