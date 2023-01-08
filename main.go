package main

import (
	"fmt"
	"go_test/page"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db     *gorm.DB
	router *gin.Engine
	log    zerolog.Logger
)

func main() {

	router := gin.Default()
	dsn := "postgres://" + "sonar" + ":" + "saidhana" + "@" + "localhost" + ":" + "5432" + "/" + "postgres"
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		fmt.Println("err", err)
		log.Error().Err(err).Msg("")
	}
	fmt.Println("db connected succesfully")
	defer db.Close()

	db.DB().SetMaxIdleConns(10)
	db.LogMode(true)
	db.SingularTable(true)

	router.Use(func(c *gin.Context) {
		c.Set("DB", db)
	})

	pageSvc := &page.HandlerService{}
	pageSvc.Bootstrap(router)

	router.Run()
}
