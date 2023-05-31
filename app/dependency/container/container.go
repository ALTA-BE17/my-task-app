package container

import (
	"time"

	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/app/config"
	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/app/database"
	bd "github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/features/book/data"
	bh "github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/features/book/handler"
	bs "github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/features/book/service"
	ud "github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/features/user/data"
	uh "github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/features/user/handler"
	us "github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/features/user/service"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var App = dig.New()

func Run() {
	// Menjalankan fungsi-fungsi yang diperlukan untuk menyediakan dependensi
	if err := App.Provide(config.InitConfig); err != nil {
		panic(err)
	}
	if err := App.Provide(database.InitDatabase); err != nil {
		panic(err)
	}
	if err := App.Provide(echo.New); err != nil {
		panic(err)
	}
	if err := App.Provide(zapLogger); err != nil {
		panic(err)
	}
	if err := RegisterDataUser(App); err != nil {
		panic(err)
	}
	if err := RegisterServiceUser(App); err != nil {
		panic(err)
	}
	if err := RegisterHandlerUser(App); err != nil {
		panic(err)
	}
	if err := RegisterDataBook(App); err != nil {
		panic(err)
	}
	if err := RegisterServiceBook(App); err != nil {
		panic(err)
	}
	if err := RegisterHandlerBook(App); err != nil {
		panic(err)
	}
}

func RegisterDataUser(c *dig.Container) error {
	if err := c.Provide(ud.New); err != nil {
		return err
	}
	return nil
}

func RegisterServiceUser(c *dig.Container) error {
	if err := c.Provide(us.New); err != nil {
		return err
	}
	return nil
}

func RegisterHandlerUser(c *dig.Container) error {
	if err := c.Provide(uh.New); err != nil {
		return err
	}
	return nil
}

func RegisterDataBook(c *dig.Container) error {
	if err := c.Provide(bd.New); err != nil {
		return err
	}
	return nil
}

func RegisterServiceBook(c *dig.Container) error {
	if err := c.Provide(bs.New); err != nil {
		return err
	}
	return nil
}

func RegisterHandlerBook(c *dig.Container) error {
	if err := c.Provide(bh.New); err != nil {
		return err
	}
	return nil
}

func zapLogger() *zap.Logger {
	config := zap.NewProductionConfig()
	timeEncoder := func(t time.Time, e zapcore.PrimitiveArrayEncoder) {
		e.AppendString(time.Now().UTC().Format("2006-01-02 15:04:05"))
	}
	config.EncoderConfig.EncodeTime = timeEncoder
	logger, err := config.Build()
	if err != nil {
		panic(err)
	}
	return logger
}
