package main
import (
  "authentication/controllers"
  "authentication/database"
  "authentication/middlewares"
  "github.com/gin-gonic/gin"
)
func main() {
  // Initialize Database
  database.Connect("root:Susan@2022@tcp(localhost:3306)/jwt_demo?parseTime=true")
  database.Migrate()
  // Initialize Router
  router := initRouter()
  router.Run(":8080")
}



func initRouter() *gin.Engine {
  router := gin.Default()
  api := router.Group("/api")
  {
    api.POST("/token", controllers.GenerateToken)
    api.POST("/user/register", controllers.RegisterUser)
    secured := api.Group("/secured").Use(middlewares.Auth())
    {
      secured.GET("/ping", controllers.Ping)
    }
  }
  return router
}