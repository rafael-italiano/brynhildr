package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/rafael-italiano/brynhildr/internal/service"
	"github.com/rafael-italiano/brynhildr/internal/web"
)

func main() {

	config, err := pgx.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(fmt.Sprintf("Failed to parse database URL: %v", err))
	}
	config.OnNotice = func(c *pgconn.PgConn, n *pgconn.Notice) {
		log.Printf("PID: %d; Message: %s\n", c.PID(), n.Message)
	}
	db := stdlib.OpenDB(*config)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	accountService := service.NewAccountService(db)
	accountHandler := web.NewAccountHandler(accountService)

	router := http.NewServeMux()

	router.HandleFunc("GET /accounts", accountHandler.GetAccounts)
	router.HandleFunc("POST /accounts", accountHandler.CreateAccount)
	router.HandleFunc("GET /accounts/{id}", accountHandler.GetAccountByID)
	router.HandleFunc("PUT /accounts/{id}", accountHandler.UpdateAccount)
	router.HandleFunc("DELETE /accounts/{id}", accountHandler.DeleteAccount)
	log.Println("Starting server...")
	http.ListenAndServe(":8080", router)

}
