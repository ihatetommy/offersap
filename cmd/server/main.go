package main

import (
	"OffersApp/internal/db"
	"OffersApp/internal/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
    db.ConnectDB()

    if err := runMigrations(); err != nil {
        log.Fatalf("Could not run migrations: %v", err)
    }

    router := gin.Default()
    
    router.GET("/hello_world", func(c *gin.Context) {
        c.String(200, "Hello, World!")
    })
    routes.AddAuthRoutes(router, db.DB)
    routes.AddUserRoutes(router, db.DB)
    routes.AddItemRoutes(router, db.DB)
    routes.AddPurchaseRoutes(router, db.DB)
    
    router.Run(":3000") 
}

func runMigrations() error {
    driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
    if err != nil {
        return err
    }

    m, err := migrate.NewWithDatabaseInstance(
        "file://migrations",
        "postgres",
        driver,
    )
    if err != nil {
        return err
    }

    if err := m.Up(); err != nil && err != migrate.ErrNoChange {
        return err
    }

    return nil
}
