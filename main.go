package main

import (
	"os"

	"github.com/ed/gaintime/config"
	"github.com/ed/gaintime/controller"
	"github.com/ed/gaintime/middleware"
	"github.com/ed/gaintime/repository.go"
	"github.com/ed/gaintime/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.DBconnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	todoRepository repository.TodoRepository = repository.NewTodoRepository(db)
	jwtService     service.JWTService        = service.NewJWTService()
	authService    service.AuthService       = service.NewAuthService(userRepository)
	userService    service.UserService       = service.NewUserService(userRepository)
	todoService    service.TodoService       = service.NewTodoService(todoRepository)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
	userController controller.UserController = controller.NewUserController(userService, jwtService)
	todoController controller.TodoController = controller.NewTodoController(todoService, jwtService)
)

func main() {
	defer config.DBdisconnection(db)

	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
	}

	todoRoutes := r.Group("api/todos", middleware.AuthorizeJWT(jwtService))
	{
		todoRoutes.GET("/", todoController.All)
		todoRoutes.POST("/", todoController.Insert)
		todoRoutes.GET("/:id", todoController.FindByID)
		todoRoutes.PUT("/:id", todoController.Update)
		todoRoutes.DELETE("/:id", todoController.Delete)
	}

	checkRoutes := r.Group("/health")
	{
		checkRoutes.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message":     "Successful Go-Lives!!!!!!",
				"Docker Test": "OK!!!!!!!",
			})
		})
	}

	port := os.Getenv("PORT")
	r.Run(":" + port)
}
