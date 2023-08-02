package controller

import (
	model "movie_api/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetAllMoviesHandler(c echo.Context) error {

	// Panggil Model untuk mendapatkan semua data film dari database
	movies, err := model.GetMovies()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	// Kirimkan data film sebagai respons ke Tampilan
	return c.JSON(http.StatusOK, movies)
}

func CreateMovieHandler(c echo.Context) error {
	// Dapatkan data film dari permintaan Tampilan
	var movie model.Movie
	if err := c.Bind(&movie); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}
	// Panggil Model untuk menyimpan data film baru ke database
	if err := model.CreateMovie(&movie); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	// Kirimkan respons ke Tampilan
	return c.JSON(http.StatusCreated, movie)
}

func UpdateMovieHandler(c echo.Context) error {
	// Dapatkan ID film dari permintaan Tampilan
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid movie ID")
	}

	// Dapatkan data film yang ingin diperbarui dari permintaan Tampilan
	var updatedMovie model.Movie
	if err := c.Bind(&updatedMovie); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}

	// Panggil Model untuk memperbarui data film berdasarkan ID
	movie, err := model.UpdateMovie(uint(id), &updatedMovie)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	// Kirimkan respons ke Tampilan
	return c.JSON(http.StatusOK, movie)
}

func DeleteMovieHandler(c echo.Context) error {
	// Dapatkan ID film dari permintaan Tampilan
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid movie ID")
	}

	// Panggil Model untuk menghapus data film berdasarkan ID
	if err := model.DeleteMovie(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	// Kirimkan respons ke Tampilan
	return c.JSON(http.StatusOK, "Movie deleted successfully")
}
