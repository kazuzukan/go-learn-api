package main

import (
	"bwa-project/handler"
	"bwa-project/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/bwa-project?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		// fatal untuk memberhentikan program kalau error
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewServices(userRepository)
	userHandler := handler.NewUserHandler(userService)
	// fmt.Println(userHandler)

	router := gin.Default()
	api := router.Group("api/v1")

	api.POST("/users", userHandler.RegisterUser)
	router.Run()

	// fmt.Println("Connection to database success")

}

// handler for gin or controller
// func Handler(c *gin.Context) {
// 	dsn := "root:@tcp(127.0.0.1:3306)/bwa-project?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		// fatal untuk memberhentikan program kalau error
// 		log.Fatal(err.Error())
// 	}

// 	fmt.Println("Connection to database success")

// 	var users []user.User
// 	db.Find(&users)

// 	c.JSON(http.StatusOK, users)
// }
