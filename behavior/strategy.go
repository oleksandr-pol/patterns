package behavior

import "fmt"

type Strategy interface {
	Execute()
}

type AdditionStrategy struct {
}

func (s *AdditionStrategy) Execute() {
	fmt.Println("addition")
}

type SubtractionStrategy struct {
}

func (s *SubtractionStrategy) Execute() {
	fmt.Println("SubtractionStrategy")
}

type Context struct {
	strategy Strategy
}

func NewContext() *Context {
	return &Context{}
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) Execute() {
	c.strategy.Execute()
}
