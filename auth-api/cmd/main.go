package main

import (
	"fmt"

	"github.com/PyMarcus/FreelaIF/auth-api/auth-api/internal/adapters/config"
)

func main() {
	fmt.Println(config.ConfigSettings.DatabaseUrl)
}