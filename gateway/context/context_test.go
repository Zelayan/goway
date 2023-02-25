package goway_context

import (
	"fmt"
	"testing"
)

func TestGoWayContext_Next(t *testing.T) {
	c := NewGoWayContext(nil, nil)
	c.filters = append(c.filters, func(c *GoWayContext) {
		fmt.Println("1 start")
		c.Next()
		fmt.Println("1 end")
	})

	c.filters = append(c.filters, func(c *GoWayContext) {
		fmt.Println("2 start")
		c.Next()
		fmt.Println("2 end")
	})

	c.Next()
}
