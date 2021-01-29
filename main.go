package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"net/http"
	"os"
)

type Connect struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func main() {
	// Echo インスタンスをNew
	e := echo.New()

	// ルーティング設定
	e.GET("/connects/:id", connectHandler)

	// サーバ起動
	e.Start(":" + os.Getenv("PORT"))
}

// ハンドラ
func connectHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, getConnect())
}

func getConnect() *Book {
	databaseUrl := os.Getenv("DATABASE_URL")
	db, err := gorm.Open("postgres", databaseUrl)
	log.Print(db)
	log.Print(err)
}
