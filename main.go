package main

import (
	"merchandise-review-list-backend/controller"
	"merchandise-review-list-backend/db"
	"merchandise-review-list-backend/repository"
	"merchandise-review-list-backend/router"
	"merchandise-review-list-backend/usecase"
	"merchandise-review-list-backend/validator"
)

func main() {
	db := db.NewDB()
	userValidator := validator.NewUserValidator()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NweUserUsecase(userRepository, userValidator)
	userController := controller.NewUserController(userUsecase)

	productRepository := repository.NewProductRepository(db)
	productUsecase := usecase.NweProductUsecase(productRepository)
	productController := controller.NewProductController(productUsecase)

	reviewPostValidator := validator.NewReviewPostValidator()
	reviewPostRepository := repository.NewPostRepository(db)
	reviewPostUsecase := usecase.NewReviewPostUsecase(reviewPostRepository, reviewPostValidator)
	reviewPostController := controller.NewReviewPostController(reviewPostUsecase)

	likeRepositor := repository.NewLikeRepository(db)
	likeUsecase := usecase.NewLikeUsecase(likeRepositor)
	likeController := controller.NewLikeController(likeUsecase)

	e := router.NewRouter(userController, productController, reviewPostController, likeController)
	e.Logger.Fatal(e.Start(":8080"))
}
