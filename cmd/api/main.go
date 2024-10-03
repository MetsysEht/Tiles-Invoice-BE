package main

import (
	"context"
	"github.com/MetsysEht/Tiles-Invoice-BE/internal/boot"
	"github.com/MetsysEht/Tiles-Invoice-BE/internal/server"
)

func main() {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	boot.Initialize()
	s := server.NewServer(boot.Config.App.Interfaces.Service)
	s.Start()
}
