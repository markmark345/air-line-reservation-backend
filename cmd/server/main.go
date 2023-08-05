package main

import (
	"air-line-reservation-backend/config"
	"air-line-reservation-backend/internal/application/services"
	"air-line-reservation-backend/internal/infrastucture/api/handler"
	"air-line-reservation-backend/internal/infrastucture/postgres"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

func main() {
	cf := config.NewConfig()
	if cf == nil {
		panic("Config is nil")
	}

	g := gin.New()
	err := registerRoutes(g, cf)
	if err != nil {
		panic(err)
	}

	startPort := fmt.Sprintf(":%d", cf.Server.Port)
	g.Run(startPort)
}

func registerRoutes(g *gin.Engine, cfg *config.Config) error {
	db, err := NewDB(cfg)
	if err != nil {
		panic(err)
	}

	userRepo := postgres.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	v1 := g.Group("/api/v1")
	{
		v1.GET("/user/:id", userHandler.GetUser)
		v1.POST("/user", userHandler.CreateUser)
	}
	return nil
}

func NewDB(cf *config.Config) (*pg.DB, error) {
	db := pg.Connect(&pg.Options{
		Addr:     cf.Postgres.Host + ":" + cf.Postgres.Port,
		User:     cf.Postgres.User,
		Password: cf.Postgres.Password,
		Database: cf.Postgres.DbName,
		PoolSize: cf.Postgres.PoolSize,
	})

	_, err := db.Exec("SELECT 1")
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
