package main

import (
	"context"
	"log"
	"user-api/db/sqlc"
	"user-api/internal/handler"
	"user-api/internal/logger"
	"user-api/internal/repository"
	"user-api/internal/routes"
	"user-api/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	logg := logger.New()
	defer logg.Sync()

	dbpool, err := pgxpool.New(context.Background(),
		"postgres://userapi:password@localhost:5432/userdb",
	)
	if err != nil {
		log.Fatal(err)
	}
	defer dbpool.Close()

	queries := sqlc.New(dbpool)
	repo := repository.NewUserRepository(queries)
	svc := service.NewUserService(repo)
	h := handler.NewUserHandler(svc)

	app := fiber.New()
	routes.SetupRoutes(app, h)

	log.Println("🚀 Server running on http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
