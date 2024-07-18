package main

import (
	"TaskManagerGoLang/config"
	"TaskManagerGoLang/controllers"
	"TaskManagerGoLang/database"
	"TaskManagerGoLang/models"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadDataBaseConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := database.NewPostgresDb(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	models.Migrate(db)

	router := gin.Default()

	router.POST("/register", func(c *gin.Context) { controllers.Register(c, db) })
	router.POST("/login", func(c *gin.Context) { controllers.Login(c, db) })

	protected := router.Group("/")
	protected.Use(controllers.AuthMiddleware(db))
	{
		protected.POST("/tasks", func(c *gin.Context) { controllers.CreateTask(c, db) })
		protected.GET("/tasks", func(c *gin.Context) { controllers.GetTasks(c, db) })
		protected.PUT("/tasks/:id", func(c *gin.Context) { controllers.UpdateTask(c, db) })
		protected.DELETE("/tasks/:id", func(c *gin.Context) { controllers.DeleteTask(c, db) })
	}

	router.Run(":8080")
}
