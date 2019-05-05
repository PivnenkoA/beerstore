package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func addAlcoholic(n, nc string) {
	var alco string = "'" + n + "'"
	var exec string = fmt.Sprintf("insert into persons (name, status, number_of_card) values (%s, 'alcoholic', %s)", alco, nc)
	db, err := sql.Open("sqlite3", "beerstore.db3")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec(exec)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
	fmt.Printf("К нам пришел еще один алкаш. Его зовут %s. Ему дали карту с номером %s\n", n, nc)
}

func main() {
	addAlcoholic("Симон Рафаэльевич Бор", "99452496")
}
