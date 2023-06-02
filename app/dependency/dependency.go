package dependency

import (
	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/app/config"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Dependency struct {
	dig.In
	Config *config.AppConfig
	Logger *zap.Logger
	DB     *gorm.DB
	Echo   *echo.Echo
}
