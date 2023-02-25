package router

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"
)

var router = make(map[string]*Router)

type Router struct {
	Path    string // 路由地址
	Url     string
	Details *routerDetails
}

type routerDetails struct {
	Path        string        // 匹配的路由地址
	Url         *url.URL      // 目标路由
	StripPrefix bool          // 是否去掉前缀
	Timeout     time.Duration // 超时时间
}

func Match(path string) (*url.URL, *Router, error) {
	route := match(path)
	if route == nil {
		return nil, nil, errors.New("no such router")
	}

	remoteUrl := &url.URL{
		Host:   route.Details.Url.Host,
		Scheme: route.Details.Url.Scheme,
		Path:   path,
	}
	if route.Details.StripPrefix {
		remoteUrl.Path = strings.TrimPrefix(path, strings.TrimRight(route.Details.Path, "**"))
	}
	return remoteUrl, route, nil
}

// match gen router
func match(path string) *Router {
	paths := strings.Split(strings.TrimLeft(path, "/"), "/")
	for index, path := range paths {
		if index == 0 {
			if r, ok := router[path]; ok {
				return r
			} else {
				return nil
			}
		}
	}
	return nil
}

func InitRouter() error {
	readRouter, err := readRouterConfig()
	if err != nil {
		return fmt.Errorf("read router config failed: %w", err)
	}
	router = readRouter
	return nil
}

// readRouterConfig
//
//	@Description: get router config
//	@return map[string]*Router
//	@return error
func readRouterConfig() (map[string]*Router, error) {
	router["hello"] = &Router{
		Path: "hello",
		Url:  "http://localhost:7070",
		Details: &routerDetails{
			Path: "/hello/**",
			Url: &url.URL{
				Scheme: "http",
				Host:   "localhost:7070",
			},
			StripPrefix: false,
			Timeout:     0,
		},
	}

	router["server1"] = &Router{
		Path: "server1",
		Url:  "http://localhost:8080",
		Details: &routerDetails{
			Path: "/server1/**",
			Url: &url.URL{
				Scheme: "http",
				Host:   "localhost:8080",
			},
			StripPrefix: false,
			Timeout:     0,
		},
	}

	router["server2"] = &Router{
		Path: "server2",
		Url:  "http://localhost:8081",
		Details: &routerDetails{
			Path: "/server2/**",
			Url: &url.URL{
				Scheme: "http",
				Host:   "localhost:8081",
			},
			StripPrefix: false,
			Timeout:     0,
		},
	}
	return router, nil
}
