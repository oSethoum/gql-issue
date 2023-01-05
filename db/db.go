package db

import (
	"api/ent"
	"api/ent/migrate"
	"context"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Client *ent.Client

func Init() {
	client, err := ent.Open("sqlite3", "file:ent.sqlite?_fk=1")

	if err != nil {
		log.Fatalln(err)
	}

	err = client.Schema.Create(
		context.Background(),
		migrate.WithDropColumn(true),
		migrate.WithDropIndex(true),
		migrate.WithGlobalUniqueID(true),
	)

	if err != nil {
		panic(err)
	}

	Client = client
}
