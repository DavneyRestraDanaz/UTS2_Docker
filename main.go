package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Struktur data untuk item yang akan disimpan
type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// Simulasi database sederhana menggunakan slice
var items = []Item{
	{ID: 1, Name: "Laptop", Price: 10000000},
	{ID: 2, Name: "Smartphone", Price: 5000000},
	{ID: 3, Name: "Tablet", Price: 3000000},
}

func main() {
	// Membuat router Gin dengan mode default
	router := gin.Default()

	// Rute untuk halaman utama
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Selamat datang di API Toko Sederhana",
		})
	})

	// Mendapatkan semua item
	router.GET("/items", func(c *gin.Context) {
		c.JSON(http.StatusOK, items)
	})

	// Mendapatkan item berdasarkan ID
	router.GET("/items/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID harus berupa angka"})
			return
		}
		for _, item := range items {
			if item.ID == id {
				c.JSON(http.StatusOK, item)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Item tidak ditemukan"})
	})

	// Menambahkan item baru
	router.POST("/items", func(c *gin.Context) {
		var newItem Item
		if err := c.ShouldBindJSON(&newItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Menentukan ID baru (biasanya dilakukan oleh database)
		newItem.ID = len(items) + 1
		items = append(items, newItem)
		c.JSON(http.StatusCreated, newItem)
	})

	// Menjalankan server pada port 8080
	router.Run(":8080")
}
