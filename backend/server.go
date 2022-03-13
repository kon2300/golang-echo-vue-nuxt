package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kon2300/go-echo-vue-nuxt/models"
	"github.com/kon2300/go-echo-vue-nuxt/routes"
)

func envLoad() {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatalf("Error loading env target")
	}
}

func main() {
	envLoad()

	db := models.Connect()

	db.AutoMigrate(&models.User{}, &models.Article{}, &models.ArticleLike{})

	routes.Init()
}
