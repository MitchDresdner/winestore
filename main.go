package main

import (
	"fmt"
	"github.com/mjd/winestore/models"
	"github.com/mjd/winestore/util"
	"log"
	"net/http"
)

type Env struct {
	db models.Datastore
}

func main () {
	fmt.Println("Starting")

	connStr, err := util.FetchYAML()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// "postgres://goland:goland@localhost/wines"
	db, err := models.Connect(connStr)
	if err != nil {
		log.Panic(err)
	}

	env := &Env{db}

	err = models.LoadDb(db)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	env.getAllWine()

	//http.HandleFunc("/books", env.booksIndex)
	//http.ListenAndServe(":3000", nil)

	//wines, err := env.db.AllWines()
	//if err != nil {
	//	log.Fatalf("error: %v", err)
	//}
	//
	//wine, e := env.db.WineById()
	//if e != nil {
	//	log.Fatalf("error: %v", e)
	//}
	//fmt.Println("You selected %s", wine.Product)

	//for _, wine := range wines {
	//	fmt.Printf("%-25s %-69s Today Only $%4.2f\n", wine.Product, wine.Description, wine.Price)
	//}

	//env.getAllWine(env)

	err = models.Close(db)

	fmt.Println("Exit")
}

func (env *Env) booksIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	wines, err := env.db.AllWines()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	for _, wine := range wines {
		fmt.Fprintf(w, "%s, %s, £%.2f\n", wine.Product, wine.Description, wine.Price)
	}
}

func (env *Env) getAllWine() {

	wines, err := env.db.AllWines()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	//Display results to console
	//fmt.Printf("%-25s %-69s Today Only $%4.2f\n",wine.product, wine.description, wine.priceprice)
	for _, wine := range wines {
		fmt.Printf("%-25s %-69s Today Only $%4.2f\n", wine.Product, wine.Description, wine.Price)
	}
}

func DeprecateMe() {

	fmt.Println("Starting")

	//models.MyModel()
	//models.A()
	//models.B()
	//models.C()
	//
	//models.D()
	//models.E()
	//models.F()
	//
	//view.MyView()
	//controller.MyController()

	fmt.Println("Exit")
}


