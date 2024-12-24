package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"url-shortener/config"
	"url-shortener/logger"
	"url-shortener/rest/controller"
	"url-shortener/rest/middleware"
	"url-shortener/rest/route"
	"url-shortener/service"
	"url-shortener/store"

	"github.com/gofiber/fiber/v2"
	"github.com/patrickmn/go-cache"
	"go.uber.org/zap"
	// "github.com/go-redis/redis/v8"
)

var (
	memcache *cache.Cache
)

// @title Url Shortener
// @version 1.0
// @description API creating and managing short links

// @BasePath /api/v1

func main() {
	ctx := context.Background()
	// config + logger
	config.Get()
	fiberConf := config.GetFiberConfig()
	l := logger.Get()

	// Init memcache
	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 10 minutes
	memcache = cache.New(5*time.Minute, 10*time.Minute)

	// Init repository store
	store, err := store.NewStore(ctx)
	if err != nil {
		l.Fatal("Failed to initialize store", zap.Error(err))
	}

	// Init service manager
	serviceManager, err := service.NewServiceManager(ctx, store)
	if err != nil {
		l.Fatal("Failed to initialize service manager", zap.Error(err))
	}

	// Init controllers
	urlsController := controller.NewUrlsController(
		ctx, serviceManager, l, memcache,
	)

	// Init fiber instance
	app := fiber.New(fiberConf)

	// Middlewares.
	middleware.FiberMiddleware(app)

	fmt.Println("Server started")

	// Routes
	route.SwaggerRoutes(app)
	route.PublicRoutes(app)
	route.PrivateRoutes(app, memcache, urlsController)

	// Start server
	log.Fatal(app.Listen(":3000"))
}
