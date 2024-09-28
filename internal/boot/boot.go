package boot

import (
	"fmt"
	"github.com/MetsysEht/Tiles-Invoice-BE/internal/config"
)

var (
	// Config contains application configuration values.
	Config config.Config
)

func init() {
	initConfig()
}

func Initialize() {
	fmt.Println("Boot Initialize")
}
