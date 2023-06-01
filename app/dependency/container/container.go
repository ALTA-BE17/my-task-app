package container

import (
	"os"

	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/app/config"
	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/app/database"
	ud "github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/features/user/data"
	uh "github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/features/user/handler"
	us "github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/features/user/service"
	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/helper"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
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

func zapLogger() *zap.Logger {
	production := false // Set this to true if running in production
	cfg := helper.ZapGetConfig(production)

	// constructing our prependEncoder with a ConsoleEncoder using your original configs
	enc := &helper.PrependEncoder{
		Encoder: zapcore.NewConsoleEncoder(cfg.EncoderConfig),
		Pool:    buffer.NewPool(),
	}

	logger := zap.New(
		zapcore.NewCore(
			enc,
			os.Stdout,
			zapcore.DebugLevel,
		),
		// this mimics the behavior of NewProductionConfig.Build
		zap.ErrorOutput(os.Stderr),
	)

	logger.Info("this is info")
	logger.Debug("this is debug")
	logger.Warn("this is warn")

	return logger
}
