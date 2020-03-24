package main

import (
	"github.com/ajordi/todo/internal/di"
	"github.com/ajordi/todo/pkg/adding"
	"github.com/ajordi/todo/pkg/authenticating"
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
	db.AutoMigrate(&authenticating.User{})

	return db
}

func main(){
	db := initDB()
	defer db.Close()

	app := di.InitApp(db)
	apiV1prefix := "/api/v1"

	r := gin.Default()
	r.POST("/auth/login", app.AuthMiddleware.LoginHandler)
	r.GET("/auth/refresh_token", app.AuthMiddleware.RefreshHandler)
	r.POST(apiV1prefix + "/users", app.Authenticating.Create)

	auth := r.Group(apiV1prefix)
	auth.Use(app.AuthMiddleware.MiddlewareFunc())
	{
		auth.POST("/tasks", app.Adding.Create)
		auth.GET("/tasks", app.Listing.FindAll)
		auth.GET("/tasks/:id", app.Listing.FindByID)
		auth.DELETE("/tasks/:id", app.Deleting.Delete)
	}

	errR := r.Run()
	if errR != nil {
		panic(errR)
	}
}
