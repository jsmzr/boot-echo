package boot

import (
	"fmt"
	"testing"

	"github.com/labstack/echo/v4"
)

type TestMiddleware struct{}
type TestErrorMiddleware struct{}

func (m *TestMiddleware) Load(e *echo.Echo) error {
	return nil
}

func (m *TestMiddleware) Order() int {
	return 0
}

func (m *TestErrorMiddleware) Load(e *echo.Echo) error {
	return fmt.Errorf("laod middleware error")
}

func (m *TestErrorMiddleware) Order() int {
	return 1
}

func TestRegisterMiddleware(t *testing.T) {
	middlewares = make(map[string]EchoMiddleware)
	name := "test"
	RegisterMiddleware(name, &TestMiddleware{})
	if _, ok := middlewares[name]; !ok {
		t.Fatal("register faield")
	}
	defer func() {
		if err := recover(); err == nil {
			t.Fatal("middleware already exists")
		}
	}()
	RegisterMiddleware(name, &TestErrorMiddleware{})
}

func TestInitMiddleware(t *testing.T) {
	middlewares = make(map[string]EchoMiddleware)
	e := echo.New()
	if err := initMiddleware(e); err != nil {
		t.Fatal(err)
	}
	RegisterMiddleware("test", &TestMiddleware{})
	if err := initMiddleware(e); err != nil {
		t.Fatal(err)
	}
	RegisterMiddleware("testError", &TestErrorMiddleware{})
	if err := initMiddleware(e); err == nil {
		t.Fatal("init should be error")
	}
}
