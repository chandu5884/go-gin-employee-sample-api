package main

import (
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    _ "github.com/mattn/go-sqlite3"
    "go-gin-api/handler"
    "go-gin-api/repository"
    "go-gin-api/service"
	"go-gin-api/model"
)

func main() {
    router := gin.Default()

    db, err := gorm.Open("sqlite3", "test.db")
    if err != nil {
        panic(err)
    }
    db.AutoMigrate(&model.Employee{})

    employeeRepo := repository.NewGormEmployeeRepository(db)
    employeeService := services.NewEmployeeService(employeeRepo)
    employeeHandler := handler.NewEmployeeHandler(employeeService)

    router.POST("/employees", employeeHandler.CreateEmployee)
    router.GET("/employees/:id", employeeHandler.GetEmployeeByID)

    router.Run(":8080")
}
