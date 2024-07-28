package main

import (
	"github.com/DominikKoniarz/go-web-files-encryptor/internal/server"
)

func main() {
	r := server.New()

	r.RegisterRoutes()

	r.Run("localhost:8080")

}
