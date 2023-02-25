package log

import (
	goway_context "github.com/Zelayan/goway/gateway/context"
	"log"
	"time"
)

func Logger() goway_context.Filter {
	return func(c *goway_context.GoWayContext) {
		t := time.Now()
		c.Next()
		log.Printf("[%d] %s in %v", c.StatusCode, c.Request.RequestURI, time.Since(t))
	}
}
