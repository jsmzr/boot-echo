package boot

import (
	"fmt"
	"sort"

	"github.com/labstack/echo/v4"
)

type EchoMiddleware interface {
	Load(*echo.Echo) error
	Order() int
}

var middlewares map[string]EchoMiddleware

func RegisterMiddleware(name string, m EchoMiddleware) {
	_, ok := middlewares[name]
	if ok {
		panic(fmt.Errorf("echo middleware [%s] already registerd", name))
	}
	log(fmt.Sprintf("Register [%s:%T] middleware", name, m))
	middlewares[name] = m
}

func initMiddleware(e *echo.Echo) error {
	if len(middlewares) == 0 {
		log("Not found Echo Middleware")
		return nil
	}
	values := make([]EchoMiddleware, 0, len(middlewares))
	for _, v := range middlewares {
		values = append(values, v)
	}
	sort.Slice(values, func(i, j int) bool {
		return values[i].Order() < values[j].Order()
	})
	for i := 0; i < len(values); i++ {
		if err := values[i].Load(e); err != nil {
			return err
		}
		log(fmt.Sprintf("Load [%T] middleware", values[i]))
	}

	return nil
}
