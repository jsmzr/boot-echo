package boot

import (
	"fmt"
	"testing"

	plugin "github.com/jsmzr/boot-plugin"
	"github.com/labstack/echo/v4"
)

type TestErrorPlugin struct {
}

func (t *TestErrorPlugin) Enabled() bool {
	return true
}

func (t *TestErrorPlugin) Order() int {
	return 1
}

func (t *TestErrorPlugin) Load() error {
	return fmt.Errorf("test error")
}

func TestRegisterRouter(t *testing.T) {
	RegisterRouter(func(e *echo.Echo) {})
}

func TestRun(t *testing.T) {
	plugin.Register("testError", &TestErrorPlugin{})
	if Run() == nil {
		t.Fatal("init plugin should be error")
	}
}
