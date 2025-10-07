package repository

import (
	"database/sql"

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
