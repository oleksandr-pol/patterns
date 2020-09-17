package structural

import (
	"fmt"
)

type OrderDetails struct {
}

func NewOrderDetails() *OrderDetails {
	return &OrderDetails{}
}

func (c *OrderDetails) SetOrderDetails() {
	fmt.Println("Order Details")
}

type OrderVolume struct {
}

func NewOrderVolume() *OrderVolume {
	return &OrderVolume{}
}

func (c *OrderVolume) SetOrderVolume() {
	fmt.Println("Order VOlume")
}

type OrderPrice struct {
}

func NewOrderPrice() *OrderPrice {
	return &OrderPrice{}
}

func (c *OrderPrice) SetOrderPrice() {
	fmt.Println("Order Price")
}

type OrderFacade struct {
	details *OrderDetails
	volume  *OrderVolume
	price   *OrderPrice
}

func NewOrderFacade() *OrderFacade {
	return &OrderFacade{
		details: NewOrderDetails(),
		volume:  NewOrderVolume(),
		price:   NewOrderPrice(),
	}
}

func (c *OrderFacade) CreateOrder() {
	c.details.SetOrderDetails()
	c.volume.SetOrderVolume()
	c.price.SetOrderPrice()
}
