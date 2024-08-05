package main

import (
	"github.com/maoustan/lotto-api/internal/api"
)

func main() {
	// db.Init()
	r := api.SetupRouter()
	r.Run(":8080")
}
