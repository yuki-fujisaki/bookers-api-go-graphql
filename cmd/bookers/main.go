package main

import (
	"bookers"
	"bookers/ent"
	"context"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-sql-driver/mysql"
)

func main() {
	entOptions := []ent.Option{}

	// 発行されるSQLをロギングするなら
	entOptions = append(entOptions, ent.Debug())

	mc := mysql.Config{
		User:                 "root",
		Passwd:               "password",
		Net:                  "tcp",
		Addr:                 "bookers-db-mysql:3306",
		DBName:               "bookers-db-mysql",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	client, err := ent.Open("mysql", mc.FormatDSN(), entOptions...)
	if err != nil {
		log.Fatalf("Error open mysql ent client: %v\n", err)
	}

	defer client.Close()

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	srv := handler.NewDefaultServer(bookers.NewSchema(client))
	http.Handle("/",
		playground.Handler("Book", "/query"),
	)
	http.Handle("/query", srv)
	log.Println("listening on :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal("http server terminated", err)
	}
}
