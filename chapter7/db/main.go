package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Person struct {
	Name string
	Age  int
}

var DBConnection *sql.DB

func main() {
	DBConnection, _ := sql.Open("sqlite3", "./example.sql")
	defer DBConnection.Close()
	cmd := `CREATE TABLE IF NOT EXISTS person(
		  name STRING,
		  age INT)
	`
	_, err := DBConnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	cmd = "DELETE FROM person"
	_, err = DBConnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	persons := []Person{{Name: "Nancy", Age: 20}, {Name: "Mike", Age: 24}, {Name: "Nina", Age: 20}}

	for _, p := range persons {
		cmd = "INSERT INTO person (name, age) VALUES (?, ?)"
		_, err = DBConnection.Exec(cmd, p.Name, p.Age)
		if err != nil {
			log.Fatalln(err)
		}
	}

	cmd = "UPDATE person SET age = ? WHERE name = ?"
	_, err = DBConnection.Exec(cmd, 25, "Mike")
	if err != nil {
		log.Fatalln(err)
	}

	// cmd = "SELECT * FROM person"
	// rows, _ := DBConnection.Query(cmd)
	// defer rows.Close()
	// var pp []Person
	// for rows.Next() {
	// 	var p Person
	// 	err := rows.Scan(&p.Name, &p.Age)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	pp = append(pp, p)
	// }
	// for _, p := range pp {
	// 	fmt.Println(p.Name, p.Age)
	// }

	cmd = "SELECT * FROM person where age = ?"
	row := DBConnection.QueryRow(cmd, 19000)

	var p Person
	err = row.Scan(&p.Name, &p.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No row")
		} else {
			log.Println(err)
		}
	}
	fmt.Println(p.Name, p.Age)
}
