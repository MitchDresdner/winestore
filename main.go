package main

import (
	"fmt"
	"github.com/mjd/winestore/models"
	"github.com/mjd/winestore/util"
	"log"
)

type Env struct {
	db models.Datastore
}

func main() {

	// 1 - Fetch database properties stored as YAML, decode secrets
	connStr, err := util.FetchYAML()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Connect to Postgres DB
	db, err := models.Connect(connStr)
	if err != nil {
		log.Panic(err)
	}

	// Save db handle for dependency injecting into supporting routines
	env := &Env{db}

	// Load some initial sample rows, truncates existing
	err = models.LoadDb(db)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Retrieve and display rows created
	env.getAllWine()

	err = models.Close(db)
}

func (env *Env) getAllWine() {

	wines, err := env.db.AllWines()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	for _, wine := range wines {
		fmt.Printf("%-25s %-69s Today Only $%4.2f\n", wine.Product, wine.Description, wine.Price)
	}
}
