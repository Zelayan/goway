package limit

import (
	"errors"
	context "github.com/Zelayan/goway/gateway/context"
	"github.com/Zelayan/goway/gateway/response"
	"github.com/juju/ratelimit"
	"log"
	"net/http"
	"time"
)

// MaxAllowed 最基本的限流 利用channel
func MaxAllowed(n int) context.Filter {
	sem := make(chan struct{}, n)
	acquire := func() {
		sem <- struct{}{}
	}
	release := func() {
		<-sem
	}
	return func(c *context.GoWayContext) {
		acquire()
		c.Next()
		release()
	}
}

// MaxAllowedWithTimeOut 带超时时间的限流，当前容量大于限制次数时，返回错误信息
func MaxAllowedWithTimeOut(max int, timeout time.Duration) context.Filter {
	sem := make(chan struct{}, max)

	return func(c *context.GoWayContext) {
		var called, fulled bool
		defer func() {
			if called == false && fulled == false {
				<-sem
			}
			if err := recover(); err != nil {
				log.Printf("err:%v", err)
			}
		}()

		select {
		case sem <- struct{}{}:
			c.Next()
			called = true
			<-sem
		case <-time.After(timeout):
			fulled = true
			log.Printf("to many request, time out")
			c.AbortWithStatus(http.StatusGatewayTimeout, response.LimitTimeOutError, errors.New(response.StatusText(response.LimitTimeOutError)))
			return
		}
	}
}

// RateLimit 基于令牌桶的限流
func RateLimit(fillInterval time.Duration, cap int) context.Filter {
	// 初始化令牌桶
	bucket := ratelimit.NewBucket(fillInterval, int64(cap))
	return func(c *context.GoWayContext) {
		if bucket.TakeAvailable(1) < 1 {
			log.Printf("to many request, time out")
			c.AbortWithStatus(http.StatusGatewayTimeout, response.LimitTimeOutError, errors.New(response.StatusText(response.LimitTimeOutError)))
			return
		}
		c.Next()
	}
}
