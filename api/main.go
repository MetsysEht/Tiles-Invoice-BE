package main

import (
	"context"
	"fmt"
	"github.com/MetsysEht/Tiles-Invoice-BE/internal/boot"
)

func main() {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	fmt.Println("Hello World")
	boot.Initialize()
}
