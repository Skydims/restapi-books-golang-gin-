package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []Book{
	{"1", "Golang Basics", "Dimas adi saputra"},
	{"2", "Advanced Golang", "Jane Doe"},
}

func getBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

func getBookByID(c *gin.Context) {
	id := c.Param("id")
	for _, book := range books {
		if book.ID == id {
			c.JSON(http.StatusOK, book)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

func addBook(c *gin.Context) {
	var newBook Book
	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	books = append(books, newBook)
	c.JSON(http.StatusCreated, newBook)
}


func main() {
    r := gin.Default()

    // Route untuk halaman utama
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Selamat datang di REST API Buku!",
        })
    })

    // Route lain yang sudah Anda miliki
    r.GET("/books", getBooks)
    r.GET("/books/:id", getBookByID)
    r.POST("/books", addBook)

    r.Run(":5000") // Menjalankan server di port 5000
}
