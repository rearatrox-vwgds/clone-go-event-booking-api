package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

// Load lädt die .env-Datei aus dem Projektroot.
// Sie kann aus jedem Service aufgerufen werden (einmal pro main.go reicht).
func Load() {
	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Println("⚠️  Keine .env-Datei gefunden – nutze System-ENV")
	}
}
