package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rahmanfadhil/gin-bookstore/config"
	"github.com/rahmanfadhil/gin-bookstore/routes"
)

func main() {

	router := gin.Default()

	// Connect to database
	config.DatabaseConnect()

	// Routes
	routes.BooksRoutes(router)

	// Run the server
	err := router.Run()
	if err != nil {
		return
	}
}
