package main

import (
	"backend/service"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	r := gin.Default()

	if port == "" {
		port = "8000"
	}

	r.POST("/login", service.LoginHandler)

	r.GET("/refresh_token", service.RefreshTokenHandler)

	userRouter := r.Group("/user")
	userRouter.Use(authMiddleware())
	userRouter.GET("/:userid/hello", service.HelloHandler)

	adminRouter := r.Group("/admin")
	adminRouter.Use(authMiddleware())
	adminRouter.GET("/hello", service.HelloHandler)

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
