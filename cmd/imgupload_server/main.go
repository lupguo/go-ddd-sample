package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/lupguo/go-ddd-sample/application"
	"github.com/lupguo/go-ddd-sample/infrastructure/persistence"
	"github.com/lupguo/go-ddd-sample/interfaces/api/handler"
	"log"
	"os"
)

func init() {
	// To load our environmental variables.
	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}
}

func main() {
	// db detail
	dbDriver := os.Getenv("DB_DRIVER")
	host := os.Getenv("DB_HOST")
	password := os.Getenv("DB_PASSWORD")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	// 初始化基础层实例 - DB实例
	persisDB, err := persistence.NewRepositories(dbDriver, user, password, port, host, dbname)
	if err != nil {
		log.Fatal(err)
	}
	defer persisDB.Close()
	// db做Migrate
	if err := persisDB.AutoMigrate(); err != nil {
		log.Fatal(err)
	}

	// 初始化应用层实例 - 上传图片应用
	uploadImgApp := application.NewUploadImgApp(persisDB.UploadImg)
	// 初始化接口层实例 - HTTP处理
	uploadImgHandler := handler.NewUploadImgHandler(uploadImgApp)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// 静态主页
	e.Static("/", "public")

	// 图片上传
	e.POST("/upload", uploadImgHandler.Save)
	e.GET("/delete/:id", uploadImgHandler.Delete)
	e.GET("/img/:id", uploadImgHandler.Get)
	e.GET("/img-list", uploadImgHandler.GetAll)

	// Start server
	e.Logger.Fatal(e.Start(os.Getenv("LISTEN_PORT")))
}
