package main

import (
	"blog-website/database"
	"blog-website/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	database.Connect()
	err := godotenv.Load()

	if err != nil{
		log.Fatal("Error loading in .env files")
	}
	port :=os.Getenv("PORT")
	app:=fiber.New()
	routes.Setup(app)
	app.Listen(":"+port)
}