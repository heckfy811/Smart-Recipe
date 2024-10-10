package api

import (
	"smart-recipe/api/handlers"

	"github.com/gin-gonic/gin"
)

// Все запросы к серваку тут
func (s *APIServer) Run() {
	app := gin.Default()

	app.GET("/", handlers.Hello) //TODO: редирект на рецепты
	app.GET("/login", handlers.Login)
	app.POST("/sign-up", handlers.SignUp)
	app.Run(s.Addr)
}

type APIServer struct {
	Addr string
}
