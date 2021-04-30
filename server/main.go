package main

import (
	"fmt"
	"net/http"

	routes "server/routes"

	helper "server/helpers"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Gin-Gonic Server")
	helper.InitDB()
	startServer()
}

func startServer() {
	router := gin.Default()

	md := cors.DefaultConfig()
	md.AllowAllOrigins = true
	md.AllowHeaders = []string{"*"}
	md.AllowMethods = []string{"*"}
	router.Use(cors.New(md))
	router.Use(static.Serve("/", static.LocalFile("App_File_Storage", true)))
	// router.Use(static.Serve("/App_File_Storage/images", static.LocalFile("App_File_Storage/images", true)))
	router.GET("/", checkStatus())
	routes.Init(router)
	s := &http.Server{
		Addr:    ":4700",
		Handler: router,
	}
	s.ListenAndServe()
}

func checkStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "Server is running successfully !!!!!")
	}
}
