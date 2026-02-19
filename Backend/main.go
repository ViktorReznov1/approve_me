package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 1. Konfigurasi Koneksi (Pastikan nama DB benar: db_approveme)
	// format: user:password@tcp(127.0.0.1:3306)/nama_db?...
	dsn := "root:@tcp(127.0.0.1:3306)/db_approveme?charset=utf8mb4&parseTime=True&loc=Local"
	
	// 2. Membuka Koneksi
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Gagal koneksi ke database:", err)
		return
	}

	// 3. Mengatasi Error "db declared and not used"
	// Kita gunakan 'db' untuk mengecek apakah koneksi benar-benar aktif
	sqlDB, _ := db.DB()
	err = sqlDB.Ping()
	if err != nil {
		fmt.Println("Database tidak merespon!")
	} else {
		fmt.Println("Koneksi Database Berhasil ke db_approveme!")
	}

	// 4. Setup Server API
	r := gin.Default()

	// Route untuk tes (Nanti istri kamu bisa nembak URL ini dari React)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "Backend ApproveMe! siap melayani istri tercinta",
		})
	})

	// 5. Jalankan Server
	r.Run(":8080") 
}