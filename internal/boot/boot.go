package boot

import (
	"fmt"
	"github.com/MetsysEht/Tiles-Invoice-BE/internal/config"
	"github.com/MetsysEht/Tiles-Invoice-BE/pkg/logger"
)

var (
	// Config contains application configuration values.
	Config config.Config
)

func init() {
	initConfig()
}

func Initialize() {
	logger.InitLogger()
	fmt.Println("Boot Initialized")
}
