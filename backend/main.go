package main

import (
	"backend/service"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(logMiddleware())
	r.POST("/login", service.LoginHandler)
	r.GET("/logout", service.LogoutHandler)

	r.GET("/refresh_token", service.RefreshTokenHandler)

	userRouter := r.Group("/user")
	userRouter.Use(authMiddleware())
	userRouter.GET("/:userid/hello", service.HelloHandler)

	adminRouter := r.Group("/admin")
	adminRouter.Use(authMiddleware())
	adminRouter.GET("/hello", service.HelloHandler)
	return r
}

func main() {
	r := setupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
