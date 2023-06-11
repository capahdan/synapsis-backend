package routes

import (
	"log"
	"synapsis-backend/controllers"
	"synapsis-backend/middlewares"
	"synapsis-backend/repositories"
	"synapsis-backend/usecases"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func init() {
	middleware.ErrJWTMissing.Code = 401
	middleware.ErrJWTMissing.Message = "Unauthorized"
}

func Init(e *echo.Echo, db *gorm.DB) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// USER

	userRepository := repositories.NewUserRepository(db)
	userUsecase := usecases.NewUserUsecase(userRepository)
	userController := controllers.NewUserController(userUsecase)

	api := e.Group("/api/v1")
	api.POST("/login", userController.UserLogin)
	api.POST("/register", userController.UserRegister)

	user := api.Group("/user")
	user.Use(middlewares.JWTMiddleware)
	user.Any("", userController.UserCredential)
	user.PATCH("/update-information", userController.UserUpdateInformation)
	user.PUT("/update-password", userController.UserUpdatePassword)
	user.PUT("/update-profile", userController.UserUpdateProfile)

	// Category
	categoryRepository := repositories.NewCategoryRepository(db)
	categoryUsecase := usecases.NewCategoryUsecase(categoryRepository)
	categoryController := controllers.NewCategoryController(categoryUsecase)

	category := api.Group("/category")
	// category.Use(middlewares.JWTMiddleware)
	category.GET("", categoryController.GetAllCategorys)
	category.GET("/:id", categoryController.GetCategoryByID)
	category.POST("", categoryController.CreateCategory)
	category.PUT("/:id", categoryController.UpdateCategory)
	category.DELETE("/:id", categoryController.DeleteCategory)

	// Product
	productRepository := repositories.NewProductRepository(db)
	productUsecase := usecases.NewProductUsecase(productRepository)
	productController := controllers.NewProductController(productUsecase)

	product := api.Group("/product")
	// product.Use(middlewares.JWTMiddleware)
	product.GET("", productController.GetAllProducts)
	product.GET("/:id", productController.GetProductByID)
	product.POST("", productController.CreateProduct)
	product.PUT("/:id", productController.UpdateProduct)
	product.DELETE("/:id", productController.DeleteProduct)

}
