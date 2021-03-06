package main

import (
	"flag"
	"log"
	"os"

	"analyze.api/app/handlers"
	_ "analyze.api/docs"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

var (
	prod = flag.Bool("prod", true, "Enable prefork in Production")
)

func setupRoutes(app *fiber.App) {
	apiGroup := app.Group("/api", logger.New())
	apiGroup.Post("/upload", handlers.FilterData)
	apiGroup.Put("/export", handlers.Export)
	apiGroup.Put("/concat", handlers.Concat)
	apiGroup.Patch("/save", handlers.SaveRange)

	fileGroup := app.Group("/file")
	fileGroup.Static("/raw", os.Getenv("RAW_DIR"))
	fileGroup.Static("/csv", os.Getenv("FILTER_DIR"))
	fileGroup.Static("/export", os.Getenv("EXPORT_DIR"))
	fileGroup.Static("/clean", os.Getenv("CLEAN_DIR"))

	app.Get("/swagger/*", swagger.HandlerDefault)
}

func NewServer() *fiber.App {
	app := fiber.New(fiber.Config{
		BodyLimit: 100 * 1024 * 1024, // 100 mb
		Prefork:   *prod,
	})
	app.Use(cors.New(cors.Config{}))

	setupRoutes(app)

	app.Static("/", os.Getenv("FRONT_DIR")) // if serve static web

	return app
}

// @title           analyze API
// @version         1.0
// @description     analyze python backend
// @termsOfService  http://swagger.io/terms/

// @tag.name         Python
// @tag.description  running python process api

// @contact.name   Wesley
// @contact.email  wesley.ch3n.0530@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3001
// @BasePath  /
func main() {
	flag.Parse()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := NewServer()
	app.Listen(":" + os.Getenv("PORT"))
}
