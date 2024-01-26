package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spacetronot-research-team/catalog-service/database"
	"github.com/spacetronot-research-team/catalog-service/internal/router"
)

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		logrus.Fatal(err)
	}

	db, err := database.InitializeDB()
	if err != nil {
		logrus.Fatal(err)
	}

	ginEngine := gin.Default()

	ginEngine.GET("/ping", ping)
	router.Add(ginEngine, db)

	if err := ginEngine.Run(); err != nil {
		logrus.Fatal(err)
	}

}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":       http.StatusOK,
		"service_name": os.Getenv("SERVICE_NAME"),
		"mode":         os.Getenv("MODE"),
	})
}
