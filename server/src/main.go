package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	controllers "server/controller"
	"server/model"
	"server/repository"
)

func main() {
	connStr := "host=db port=5432 user=user password=password dbname=expenses sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	tagDbAdapter := repository.CreateDbAdapter(db, model.CreateTagFactory())
	tagRepository := repository.CreateTagRepository(tagDbAdapter)
	tagsController := controllers.CreateTagController(tagRepository)

	expenseDbAdapter := repository.CreateDbAdapter(db, model.CreateExpenseFactory())
	expenseRepository := repository.CreateExpenseRepository(expenseDbAdapter)
	expenseController := controllers.CreateExpenseController(expenseRepository)

	router := mux.NewRouter()

	router.HandleFunc("/api/expenses", expenseController.GetAllExpenses).Methods("GET")
	router.HandleFunc("/api/expenses", expenseController.CreateExpense).Methods("POST")

	router.HandleFunc("/api/tags", tagsController.GetAllTags).Methods("GET")
	router.HandleFunc("/api/tags", tagsController.CreateTag).Methods("POST")

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
