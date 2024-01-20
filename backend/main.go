package main

import (
	"fmt"
	"goblog/models"
	"goblog/pg_connect"
	"goblog/templater"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // Update this with the actual origin(s) of your frontend
	router.Use(cors.New(config))

	router.GET("/articles", func(c *gin.Context) {
		articles, err := pg_connect.GetAllData()
		if err != nil {
			c.String(http.StatusNotFound, "Article not found")
			return
		}
		c.JSON(200, articles)
	})

	router.GET("/articles/:index", func(c *gin.Context) {
		indexNo := c.Param("index")

		uIntValue, err := strconv.ParseUint(indexNo, 10, 64)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		var indexNoConv uint = uint(uIntValue)

		data, err := pg_connect.GetData(indexNoConv)
		if err != nil {
			c.String(http.StatusNotFound, "Article not found")
			return
		}
		c.JSON(200, data)
	})

	router.Run(":8080")

}

func loadExistingData() {
	templater.Run()
}

func loadDatabase() {
	pg_connect.Connect()
	pg_connect.Database.AutoMigrate(&models.Article{})
	loadExistingData()
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
