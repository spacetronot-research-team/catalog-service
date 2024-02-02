package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spacetronot-research-team/catalog-service/database"
	"github.com/spacetronot-research-team/catalog-service/internal/router"
	logrushook "github.com/spacetronot-research-team/catalog-service/pkg/logrus_hook"
)

func main() {
	logrus.AddHook(&logrushook.Trace{})

	if err := godotenv.Load(".env"); err != nil {
		logrus.Fatal(err)
	}

	db, err := database.InitializeDB()
	if err != nil {
		logrus.Fatal(err)
	}

	ginEngine := gin.Default()

	router.Add(ginEngine, db)

	if err := ginEngine.Run(); err != nil {
		logrus.Fatal(err)
	}
}
