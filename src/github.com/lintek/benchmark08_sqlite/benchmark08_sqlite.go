package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

const (
	NUM_ITER = 500
	NUM_RECORDS = 20
)

func initDatabase(databaseName string) *sql.DB {
	db, err := sql.Open("sqlite3", "./my_database")
	if err != nil {
		log.Fatal(err)
	}

	sqlStmt := `CREATE TABLE IF NOT EXISTS MY_DATABASE (PNAME TEXT PRIMARY KEY, PVALUE TEXT); DELETE FROM MY_DATABASE`
	_, err = db.Exec(sqlStmt)

	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		db.Close()
		return nil
	}

	return db
}

func selectFromDatabase(databaseName string, db *sql.DB) map[string]string {
	result := make(map[string]string)

	rows, err := db.Query(`SELECT PNAME, PVALUE FROM MY_DATABASE`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var key string
		var value string
		rows.Scan(&key, &value)
		result[key] = value

		// fmt.Println(key, value)
	}

	return result
}

func insertIntoDatabase(databaseName string, db *sql.DB, values map[string]map[string]string) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare(`INSERT OR REPLACE INTO MY_DATABASE VALUES(?,?)`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for _, myValue := range values {
		for sKey, sValue := range myValue {
			_, err = stmt.Exec(sKey, sValue)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	tx.Commit()
}

func closeDatabase(db *sql.DB) {
	db.Close()
}


func main() {
	fmt.Printf("Running: %s\n", os.Args[0])

	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Printf("Number of available CPUs: %d\n", runtime.NumCPU())

	// Prepare test data
	myData := make(map[string]map[string]string)

	for i := 0; i < NUM_ITER; i++ {
		mValue := make(map[string]string)
		for j := 0; j < NUM_RECORDS; j++ {
			mValue[ fmt.Sprintf(" %04d %04d", i, j)] = fmt.Sprintf(" Record value %04d %04d\n", i, j)
		}

		myData[fmt.Sprintf("%04d ", i)] = mValue
	}

	// fmt.Printf("My data: %v\n", my_data)

	// Remove old database file
	if _, err := os.Stat("my_database"); err == nil {
		err := os.Remove("my_database")
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// Initialize database
	myDB := initDatabase("my_database")

	start_time := time.Now()

	// Insert data into database
	insertIntoDatabase("my_database", myDB, myData)

	// Read data from the database
	retrievedData := selectFromDatabase("my_database", myDB)

	execution_time := time.Now().Sub(start_time)
	fmt.Printf("Execution time: %f\n", execution_time.Seconds())

	fmt.Printf("Num retrieved items: %d\n", len(retrievedData))
	//fmt.Printf("Retrieved data: %v\n", retrievedData)

	closeDatabase(myDB)
}
