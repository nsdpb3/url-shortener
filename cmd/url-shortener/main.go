package main

import (
	"fmt"
	"url-shortener/internal/config"
)

func main() {
	//TODO: init config: cleanenv
	cfg := config.MustLoad()
	fmt.Println(cfg)
	//TODO: init logger: slog
	//TODO: init storeage: sqlite
	//TODO: init router: chi, render
	//TODO: init server:
}
