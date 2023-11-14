package main

import (
	"go-backend/routers"
)

func main() {
	router := routers.SetupRouter()

	router.Run(":8080")
}
