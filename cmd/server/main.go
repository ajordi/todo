package main

import (
	"github.com/ajordi/todo/internal/di"
	"github.com/ajordi/todo/pkg/adding"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

func initDB() *gorm.DB{
	db, err := gorm.Open("mysql", os.Getenv("DB_URL"))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&adding.Task{})

	return db
}

func main(){
	db := initDB()
	defer db.Close()

	taskAPI := di.InitTaskAPI(db)

	r := gin.Default()
	r.POST("/tasks", taskAPI.Create)

	err := r.Run()
	if err != nil {
		panic(err)
	}
}
