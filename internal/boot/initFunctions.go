package boot

import (
	"github.com/MetsysEht/Tiles-Invoice-BE/pkg/config"
	"github.com/MetsysEht/Tiles-Invoice-BE/utils/osUtils"
	"log"
)

func initConfig() {
	// Init config
	err := config.NewDefaultConfig().Load(osUtils.GetEnv(), &Config)
	if err != nil {
		log.Fatal(err)
	}
}
