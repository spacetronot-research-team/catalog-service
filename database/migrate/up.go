package main

import (
	"fmt"
	"path/filepath"

	"github.com/joho/godotenv"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/sirupsen/logrus"
	"github.com/spacetronot-research-team/catalog-service/database"
	"gorm.io/gorm"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		logrus.Fatal(err)
	}

	db, err := database.InitializeDB()
	if err != nil {
		logrus.Fatal(err)
	}

	if err := migrationUp(db); err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("migrate up success")
}

func migrationUp(db *gorm.DB) error {
	migrate.SetTable("migrations")

	sql, err := db.DB()
	if err != nil {
		return fmt.Errorf("`*gorm.DB` fail return `*sql.DB`: %v", err)
	}

	_, err = migrate.Exec(sql, "postgres", getFileMigrationSource(), migrate.Up)
	if err != nil {
		return fmt.Errorf("fail execute migrations: %v", err)
	}

	return nil
}

func getFileMigrationSource() *migrate.FileMigrationSource {
	migrations := &migrate.FileMigrationSource{
		Dir: filepath.Join("database", "schema_migration"),
	}
	return migrations
}
