package models

import "fmt"

type Wine struct {
	Product string
	Description string
	Price float32
}

func (db *DB) AllWines() ([]*Wine, error) {
	// Query and display persisted posted
	rows, err := db.Query("SELECT product, pdesc, price from winetbl")
	if err != nil {
		return nil, err
	}

	err = fmt.Errorf("Query returns no data!")
	if rows == nil {
		return nil, err
	}

	wines := make([]*Wine, 0)

	fmt.Println("--- Wine Collection")
	defer  rows.Close()
	for rows.Next(){

		wine := new(Wine)

		err = rows.Scan(&wine.Product, &wine.Description, &wine.Price)
		if err != nil {
			fmt.Print(err)
		}

		wines = append(wines, wine)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return wines, nil
}


