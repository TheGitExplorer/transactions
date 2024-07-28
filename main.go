// main.go
package main

import (
	"github.com/TheGitExplorer/transactions/config"
	"github.com/TheGitExplorer/transactions/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	defer config.DB.Close()

	router := gin.Default()
	routes.RegisterTransactionRoutes(router)
	router.Run(":8080")
}
