package boot

import (
	"fmt"
	"github.com/MetsysEht/Tiles-Invoice-BE/internal/config"
	"github.com/MetsysEht/Tiles-Invoice-BE/pkg/gormDatabase"
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
	InitDatabase()

	fmt.Println("Boot Initialized")
}

func InitDatabase() {
	db, err := gormDatabase.CreateGormDatabase(&Config.DB)
	if err != nil {
		logger.Sl.Fatal("Could not connect to DB")
	}
	DB = db
}
