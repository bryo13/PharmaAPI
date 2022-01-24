package handlerStuff

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"phamacy/dbStuff"
)

// Book struct (Model)
type Pharmacy struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Stock int    `json:"stock"`
	Sales string `json:"sales"`
}

// Get all drugs
func GetDrugs(w http.ResponseWriter, r *http.Request) {
	var pharmacy []*Pharmacy
	db := dbStuff.Connect()
	w.Header().Set("Content-Type", "application/json")
	rows, err := db.Query("SELECT * FROM drugs")
	if err != nil {
		// handle this error better than this
		log.Println(err)
		return
	}
	defer rows.Close()
	for rows.Next() {

		p := new(Pharmacy)
		switch err = rows.Scan(&p.ID, &p.Name, &p.Stock, &p.Sales); err {
		case nil:
			pharmacy = append(pharmacy, p)
		default:
			log.Println(err)
		}
	}
	json.NewEncoder(w).Encode(pharmacy)
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}

}

// Get single drug
func GetDrug(w http.ResponseWriter, r *http.Request) {
	db := dbStuff.Connect()
	w.Header().Set("Content-Type", "application/json")

	sqlStatement := `SELECT * FROM drugs WHERE id=$1;`
	var pharmacy []*Pharmacy
	row := db.QueryRow(sqlStatement, 2)
	p := new(Pharmacy)
	switch err := row.Scan(&p.ID, &p.Name, &p.Stock, &p.Sales); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		pharmacy = append(pharmacy, p)
		json.NewEncoder(w).Encode(pharmacy)
	default:
		log.Println(err)
	}
}

func CreateDrug(w http.ResponseWriter, r *http.Request) {
	db := dbStuff.Connect()
	sqlStatement := `
INSERT INTO drugs (id,name,stock,sales)
VALUES ($1, $2, $3, $4)
RETURNING id`
	id := 0
	err := db.QueryRow(sqlStatement, 4, "Asprin", 200, "Jonathan").Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)

}

func UpdateDrug(w http.ResponseWriter, r *http.Request) {
	db := dbStuff.Connect()
	sqlStatement := `
UPDATE drugs
SET name = $2
WHERE id = $1;`
	_, err := db.Exec(sqlStatement, 1, "NewFirst")
	if err != nil {
		panic(err)
	}

	fmt.Println("Updated")

}

func DeleteDrug(w http.ResponseWriter, r *http.Request) {
	db := dbStuff.Connect()
	sqlStatement := `
DELETE FROM drugs
WHERE id = $1;`
	_, err := db.Exec(sqlStatement, 4)
	if err != nil {
		panic(err)
	}

	fmt.Println("deleted")

}
