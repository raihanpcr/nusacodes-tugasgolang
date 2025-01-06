package concurrency

import (
	"fmt"
	"sync"
)

//struct calculator
type Calculator struct {
	Result int
	Mutex  sync.Mutex
}

// func Add
func (c *Calculator) Add(value1, value2 int)  {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	c.Result = value1 + value2
	fmt.Println(c.Result)
}

// func Subtract
func (c *Calculator) Subtract(value1, value2 int)  {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	c.Result = value1 - value2
	fmt.Println(c.Result)
}

// func Multiply
func (c *Calculator) Multiply(value1, value2 int)  {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	c.Result = value1 * value2
	fmt.Println(c.Result)
}