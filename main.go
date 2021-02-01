package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"github.com/lib/pq"
	"log"
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
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "はろわ")
	})
	e.GET("/connects/:id", connectHandler)

	// サーバ起動
	e.Start(":" + os.Getenv("PORT"))
}

// ハンドラ
func connectHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, getConnect())
}

func getConnect() *Connect {
	databaseUrl := os.Getenv("DATABASE_URL")
	connection, err := pq.ParseURL(databaseUrl)
	connection += " sslmode=require"
	db, err := gorm.Open("postgres", connection)
	log.Print(db)
	log.Print(err)
	if err != nil {
		c := &Connect{
			Title:  "しっぱい",
			Author: "しっぱい",
		}
		return c
	} else {
		c := &Connect{
			Title:  "せいこう",
			Author: "せいこう",
		}
		return c
	}
}
