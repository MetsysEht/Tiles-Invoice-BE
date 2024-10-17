package boot

import (
	"github.com/MetsysEht/Tiles-Invoice-BE/internal/config"
	"github.com/MetsysEht/Tiles-Invoice-BE/internal/database"
	"github.com/MetsysEht/Tiles-Invoice-BE/pkg/logger"
	"gorm.io/gorm"
)

var (
	// Config contains application configuration values.
	Config config.Config
	DB     *gorm.DB
)

func init() {
	initConfig()
}

func Initialize() {
	logger.InitLogger()
	db, err := database.GetDatabase(&Config.DB)
	if err != nil {
		panic(err.Error())
	}
	DB = db
}
