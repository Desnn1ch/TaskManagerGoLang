package handlers

import (
	"TaskManagerGoLang/controllers"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

func SetupRouter(db *sqlx.DB, rdb *redis.Client) *gin.Engine {
	router := gin.Default()

	controllers.InitializeRedis(rdb)

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

	return router
}
