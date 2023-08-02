package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Title       string
	Description string
	Rating      float64
}

var db *gorm.DB

// Inisialisasi koneksi database MySQL
func InitDB(dsn string) error {
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	// AutoMigrate untuk membuat tabel movies jika belum ada
	db.AutoMigrate(&Movie{})
	return nil
}

// Fungsi CRUD untuk mengelola data film di MySQL database

// GetMovies mengambil semua data film dari database
func GetMovies() ([]Movie, error) {
	var movies []Movie
	if err := db.Find(&movies).Error; err != nil {
		return nil, err
	}
	return movies, nil
}

// CreateMovie menyimpan data film baru ke database
func CreateMovie(movie *Movie) error {
	if err := db.Create(movie).Error; err != nil {
		return err
	}
	return nil
}

// UpdateMovie memperbarui data film berdasarkan ID
func UpdateMovie(id uint, updatedMovie *Movie) (*Movie, error) {
	var movie Movie
	if err := db.First(&movie, id).Error; err != nil {
		return nil, err
	}
	movie.Title = updatedMovie.Title
	movie.Description = updatedMovie.Description
	movie.Rating = updatedMovie.Rating
	if err := db.Save(&movie).Error; err != nil {
		return nil, err
	}
	return &movie, nil
}

// DeleteMovie menghapus data film berdasarkan ID
func DeleteMovie(id uint) error {
	if err := db.Delete(&Movie{}, id).Error; err != nil {
		return err
	}
	return nil
}
