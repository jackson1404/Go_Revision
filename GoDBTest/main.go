package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// Struct mapping to DB table
type Album struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var db *sql.DB

func getAlbums() ([]Album, error) {
	rows, err := db.Query("SELECT album_id, title, artist, price FROM tbl_albums")
	if err != nil {
		return nil, err
	}
	defer rows.Close() // release the connection back to pool

	var albums []Album
	for rows.Next() {
		var a Album
		// Scan writes column values into struct fields
		if err := rows.Scan(&a.ID, &a.Title, &a.Artist, &a.Price); err != nil {
			return nil, err
		}
		albums = append(albums, a)
	}

	// check if iteration had any errors
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return albums, nil
}

func addNewAlbum(a Album) (int, error) {

	stmt, err := db.Prepare(`
					INSERT INTO tbl_albums (title, artist, price)
					VALUEs ($1, $2, $3)
					RETURNING album_id
				`)

	if err != nil {
		return 0, fmt.Errorf("prepare insert failed, %v", err)
	}
	defer stmt.Close()

	var id int
	err = stmt.QueryRow(a.Title, a.Artist, a.Price).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("execute insert failed %v", err)
	}

	return id, nil

}

// Connect to DB
func main() {

	var albums []Album
	var err error
	connStr := "user=postgres password=root dbname=go_albums_db sslmode=disable"
	db, err = sql.Open("postgres", connStr) //doesnt actually open a real net connection, just validate DB handle and configure pool.
	if err != nil {
		log.Fatal("Failed to connect:", err)
	}
	defer db.Close()

	// Check if DB is alive

	if err = db.Ping(); err != nil {
		log.Fatal("Cannot ping DB:", err)
	}
	fmt.Println("✅ Connected to PostgreSQL!")

	albums, err = getAlbums()
	if err != nil {
		log.Fatal("Query failed", err)
	}

	fmt.Println("🎵 Albums in DB:")
	for _, value := range albums {
		fmt.Printf("%d | %s | % s | %f ", value.ID, value.Artist, value.Title, value.Price)
	}

	newAlbum := Album{
		Title:  "Kind of Blue",
		Artist: "Miles Davis",
		Price:  42.50,
	}

	id, err := addNewAlbum(newAlbum)
	if err != nil {
		log.Fatal("failed to insert new album", err)
	}
	fmt.Printf("🎵 Inserted new album with ID: %d\n", id)

}
