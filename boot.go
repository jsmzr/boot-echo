package boot

import (
	"fmt"
	"time"

	plugin "github.com/jsmzr/boot-plugin"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

const configPrefix = "boot.echo."

var defaultDefault map[string]interface{} = map[string]interface{}{
	"port": 8080,
}

var routerInitFunctions []func(*echo.Echo)

func RegisterRouter(f func(*echo.Echo)) {
	routerInitFunctions = append(routerInitFunctions, f)
}

func Run() error {
	if err := plugin.PostProccess(); err != nil {
		return err
	}
	e := echo.New()
	log("init echo middleware")
	if err := initMiddleware(e); err != nil {
		return err
	}
	log("init echo router")
	for _, f := range routerInitFunctions {
		f(e)
	}
	log("inti echo service")
	return e.Start(fmt.Sprintf(":%d", viper.GetInt(configPrefix+"port")))
}

func log(message string) {
	fmt.Printf("[BOOT-ECHO] %v| %s\n", time.Now().Format("2006-01-02 15:04:05"), message)
}

func init() {
	for key := range defaultDefault {
		viper.SetDefault(configPrefix+key, defaultDefault[key])
	}
}
