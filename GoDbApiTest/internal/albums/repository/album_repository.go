package repository

import (
	"database/sql"
	"fmt"

	albumModel "jackson.com/goApiDb/internal/albums/models"
)

type AlbumRepository struct {
	DB *sql.DB
}

func NewAlbumRepository(db *sql.DB) *AlbumRepository {
	return &AlbumRepository{DB: db}
}

func (r *AlbumRepository) GetAlbums() ([]albumModel.Album, error) {

	rows, err := r.DB.Query("SELECT album_id, title, artist, price from tbl_albums")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var albums []albumModel.Album
	for rows.Next() {
		var a albumModel.Album
		if err := rows.Scan(&a.AlbumID, &a.Title, &a.Artist, &a.Price); err != nil {
			return nil, err
		}
		albums = append(albums, a)
	}

	return albums, nil
}

func (r *AlbumRepository) InsertAlbum(a albumModel.Album) (int, error) {

	stmt := `
		INSERT INTO tbl_albums (title, artist, price) 
		VALUES ($1, $2, $3)
		RETURNING album_id
	`

	err := r.DB.QueryRow(stmt, a.Title, a.Artist, a.Price).Scan(&a.AlbumID)
	if err != nil {
		return 0, err
	}
	return a.AlbumID, nil
}

func (r *AlbumRepository) UpdateAlbum(a albumModel.Album) error {
	// Step 1: Check existence
	var exists bool
	// EXISTS keyword checks whether the inner query returns at least one row.
	err := r.DB.QueryRow("SELECT EXISTS (SELECT 1 FROM tbl_albums WHERE album_id = $1)", a.AlbumID).Scan(&exists)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("album with id %d not found", a.AlbumID)
	}

	// Step 2: Perform update
	stmt := `
        UPDATE tbl_albums
        SET title = $1, artist = $2, price = $3
        WHERE album_id = $4
    `
	_, err = r.DB.Exec(stmt, a.Title, a.Artist, a.Price, a.AlbumID)
	return err
}

// Delete an album by ID
func (r *AlbumRepository) DeleteAlbum(id int) error {
	stmt := `DELETE FROM tbl_albums WHERE album_id = $1`
	result, err := r.DB.Exec(stmt, id)

	if err != nil {
		return err
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowAffected == 0 {
		return fmt.Errorf("album with id %d not found", id)
	}

	return nil
}
