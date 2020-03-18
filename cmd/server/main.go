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
	r.POST("/tasks", taskAPI.Adding.Create)
	r.GET("/tasks", taskAPI.Listing.FindAll)
	r.GET("/tasks/:id", taskAPI.Listing.FindByID)
	r.DELETE("/tasks/:id", taskAPI.Deleting.Delete)

	err := r.Run()
	if err != nil {
		panic(err)
	}
}
