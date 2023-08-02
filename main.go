package main

import (
	"fmt"
	"log"
	controller "movie_api/controllers"
	model "movie_api/models"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func main() {
	loadEnv()
	// Inisialisasi koneksi database MySQL
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))
	if err := model.InitDB(dsn); err != nil {
		panic(err)
	}

	e := echo.New()

	// Tambahkan endpoint untuk GET semua film
	e.GET("/movies", controller.GetAllMoviesHandler)

	// Tambahkan endpoint untuk POST film baru
	e.POST("/movie", controller.CreateMovieHandler)

	// Tambahkan endpoint untuk PUT film berdasarkan ID
	e.PUT("/movie/:id", controller.UpdateMovieHandler)

	// Tambahkan endpoint untuk DELETE film berdasarkan ID
	e.DELETE("/movie/:id", controller.DeleteMovieHandler)

	e.Start(":8080")
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
