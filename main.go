package main

import (
	"bwa-project/routes"
)

func main() {
	router := routes.SetupRoutes()
	router.Run()
}
