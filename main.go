package main

import (
	"database/sql"
	"go-gin-review/internal/controllers"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var router = gin.Default()
	var db *sql.DB
	var e error

	if db, e = sql.Open("sqlite3", "./data.db"); e != nil {
		log.Fatalln("Error: &v", e)
	}
	defer db.Close()

	if e := db.Ping(); e != nil {
		log.Fatalf("Error: %v", e)
	}
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS products (id INTEGER PRIMARY KEY AUTOINCREMENT,guid VARCHAR(55) UNIQUE NOT NULL,name VARCHAR(255) UNIQUE NOT NULL,price REAL NOT NULL,description TEXT,createdAt TEXT NOT NULL)"); err != nil {
		log.Fatalln("Error: &v", e)
	}

	router.GET("/products", controllers.GetProducts(db))
	router.GET("/products/:guid", controllers.GetProduct(db))
	router.POST("/products", controllers.PostProducts(db))
	router.PUT("/products/:guid", controllers.PutProducts(db))
	router.DELETE("/products/:guid", controllers.DeleteProducts(db))

	log.Fatalln(router.Run(":8000"))

}
