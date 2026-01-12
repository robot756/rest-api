package main

import (
	"fmt"
	"rest-api/internal/config"
)

func main() {
	// TODO: init config: cleanevn
	cfg := config.MustLoad()
	fmt.Println(cfg)

	// TODO: init logger: slog

	// TODO: init storage: postreSQL

	// TODO: init router: chi

	// TODO: init run server:
}
