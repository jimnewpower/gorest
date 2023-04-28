package main

import (
    "context"
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/lib/pq"

	"github.com/aws/aws-lambda-go/lambda"
)

type Vessel struct {
    ID        int
    Name      string
    Longitude float64
    Latitude  float64
    Status    string
}

type MyEvent struct {
	Name string `json:"name"`
}

func query() (string) {
    // Open a connection to the database
    // Get environment variables
    dbHost := os.Getenv("HOST") // prima.cvrj95nytzmd.us-west-2.rds.amazonaws.com
    dbPort := os.Getenv("PORT") // 5432
//    dbName := os.Getenv("") // prima
    dbUser := os.Getenv("USER") // postgres
    dbPass := os.Getenv("PASS") // TrHa0C0a3PoQSXAd0OPS

    // Open a connection to the database
	connect := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=require", dbHost, dbPort, dbUser, dbPass)
	log.Printf("Connecting to %s", connect)

    db, err := sql.Open("postgres", connect)
    if err != nil {
        log.Fatal("Failed to open DB connection: ", err)
    }
    defer db.Close()

    // Prepare the query statement
    query := "SELECT id, name, longitude, latitude, status FROM ships"
    stmt, err := db.Prepare(query)
    if err != nil {
        log.Fatal("Failed to prepare query statement: ", err)
    }
    defer stmt.Close()

    // Execute the query and process the results
    rows, err := stmt.Query()
    if err != nil {
        log.Fatal("Failed to execute query: ", err)
    }
    defer rows.Close()

    vessels := []Vessel{}

    for rows.Next() {
        // Process each row of data
        var v Vessel
        if err := rows.Scan(&v.ID, &v.Name, &v.Longitude, &v.Latitude, &v.Status); err != nil {
            log.Fatal("Failed to scan row: ", err)
        }
        vessels = append(vessels, v)
    }
    if err := rows.Err(); err != nil {
        log.Fatal("Failed to process rows: ", err)
    }

    // Print the results
	results := ""
    for _, v := range vessels {
        results += fmt.Sprintf("ID: %d, Name: %s, Longitude: %f, Latitude: %f, Status: %s\n", v.ID, v.Name, v.Longitude, v.Latitude, v.Status)
    }

	return results
}

func handleRequest(ctx context.Context, name MyEvent) (string, error) {
	return fmt.Sprintf("Query complete:\n %s", query()), nil
	// return fmt.Sprintf("Query complete (%s)", name.Name ), nil
}

func main() {
	lambda.Start(handleRequest)
}