package models

import "fmt"

func LoadDb(db *DB) error {

	wines:= [] Wine {
		{
			Product: "SOMMELIER SELECT",
			Description: "Old vine Cabernet Sauvignon",
			Price: 159.99,
		},
		{
			Product: "MASTER VINTNER",
			Description: "Pinot Noir captures luscious aromas",
			Price: 89.99,
		},
		{
			Product: "WINEMAKER'S RESERVE",
			Description: "Merlot featuring complex flavors of cherry",
			Price: 84.99,
		},
		{
			Product: "ITALIAN SANGIOVESE",
			Description: "Sangiovese grape is famous for its dry, bright cherry character",
			Price: 147.99,
		},
	}

	//check postgres to see if table exists
	var checkDatabase string
	err := db.QueryRow("SELECT to_regclass('public.winetbl')").Scan(&checkDatabase)
	if err != nil {
		return  err
	}

	// Create table if DNE
	if checkDatabase == "" {
		createSQL := "CREATE TABLE public.winetbl (id SERIAL PRIMARY KEY,product character varying,pdesc character varying,price decimal);"
		if _, err = db.Query(createSQL); err != nil {
			panic(err)
		}
		fmt.Println("Wine Database Created")
	}

	// sql to insert wine data
	statement :="INSERT INTO winetbl(product, pdesc, price) VALUES($1, $2, $3)"

	// Create prepared statement for inserts
	stmt , err := db.Prepare(statement)
	if err != nil {
		return  err
	}
	defer stmt.Close()

	// Start with a clean slate
	if _, err = db.Exec(`TRUNCATE TABLE winetbl`); err != nil {
		panic(err)
	}

	// Insert static entries into database
	for idx := 0; idx < len(wines); idx++ {
		w := wines[idx]

		// Invoke DB prepared statement
		//var theRow = &sql.Row{}
		//theRow = stmt.QueryRow(w.product, w.description, w.price)
		//fmt.Print(theRow)
		_, err := stmt.Exec(w.Product, w.Description, w.Price)
		if err != nil {
			return  err
		}

		// Exec returns meta data, which may be used as below
		//lastId, err := res.LastInsertId()
		//if err != nil {
		//	log.Fatal(err)
		//}
		//rowCnt, err := res.RowsAffected()
		//if err != nil {
		//	log.Fatal(err)
		//}

	}

	if err = stmt.Close(); err != nil {
		panic(err)
	}

	return nil
}