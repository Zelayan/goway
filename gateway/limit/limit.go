package limit

import (
	goway_context "github.com/Zelayan/goway/gateway/context"
)

// MaxAllowed 最基本的限流 利用channel
func MaxAllowed(n int) goway_context.Filter {
	sem := make(chan struct{}, n)
	acquire := func() {
		sem <- struct{}{}
	}
	release := func() {
		<-sem
	}
	return func(c *goway_context.GoWayContext) {
		acquire()
		c.Next()
		release()
	}
}
