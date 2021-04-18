package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/gorilla/mux"
)

func main() {
	pgHost := getEnv("MY_DATABASE_SERVICE_HOST", "localhost")
	pgPort := getEnv("DB_PORT", "5432")
	pgUser := getEnv("DB_USERNAME", "postgres")
	pgDBname := getEnv("DB_DBNAME", "postgres")
	pgPassword := getEnv("DB_PASSWORD", "mysecretpassword")
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		pgHost,
		pgPort,
		pgUser,
		pgDBname,
		pgPassword)
	db, err := NewFruitsRepository("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.AutoMigrate(&Fruit{})
	
	// Delete all data in Fruit table
	db.Delete(&Fruit{})

	// Sample data
	_, err  = db.CreateFruit(Fruit{Name: "Apple", Stock: 10})
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.CreateFruit(Fruit{Name:"Orange", Stock:10 })
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.CreateFruit(Fruit{Name: "Pear", Stock: 10})
	if err != nil {
		log.Fatal(err)
	}
	fruitController := NewFruitController(db)

	r := mux.NewRouter()
	r.HandleFunc("/api/fruits", fruitController.List).Methods("GET")
	r.HandleFunc("/api/fruits/{id:[0-9]+}", fruitController.Show).Methods("GET")
	r.HandleFunc("/api/fruits", fruitController.Create).Methods("POST")
	r.HandleFunc("/api/fruits/{id:[0-9]+}", fruitController.Update).Methods("PUT")
	r.HandleFunc("/api/fruits/{id:[0-9]+}", fruitController.Delete).Methods("DELETE")
	r.PathPrefix("/").Handler(http.FileServer(assetFS()))

	http.Handle("/", r)
	r.Use(loggingMiddleware)
	fmt.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("\n[%s] %q %q",
			time.Now().Format("02/Jan/2006:15:04:05 -0700"),
			r.Method,
			r.RequestURI)
		next.ServeHTTP(w, r)

	})
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
